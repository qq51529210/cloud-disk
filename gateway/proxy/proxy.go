package proxy

import (
	"gateway/cfg"
	"gateway/db"
	"net/http"
	"time"

	"github.com/qq51529210/log"
	"github.com/qq51529210/uuid"
)

// Serve 开始服务
func Serve() error {
	// 加载数据
	err := reloadDB()
	if err != nil {
		return err
	}
	// 监听
	return http.ListenAndServe(cfg.Cfg.ProxyAddr, http.HandlerFunc(handle))
}

// handle 处理所有请求
func handle(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	defer func() {
		// 异常
		log.Recover(recover())
		// 花费时间
		log.Infof("cost %v", time.Since(now))
	}()
	// 获取服务
	ser := getServer(r)
	if ser == nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer ser.Minus()
	// 访问控制
	if ser.Limite > 0 && !limite(w, r, ser.Limite) {
		return
	}
	// 身份验证
	if ser.Auth && !authorize(w, r) {
		return
	}
	// 设置头
	setHeaders(r)
	// 转发
	ser.ServeHTTP(w, r)
}

// getServer 返回代理的服务
func getServer(r *http.Request) *server {
	// 找出第一层
	path := r.URL.Path
	for i := 1; i < len(path); i++ {
		if i == '/' {
			path = path[:i]
			// 减去第一层
			r.URL.Path = r.URL.Path[i:]
			break
		}
	}
	// 获取服务
	return getMinLoadServer(path)
}

// setHeaders 设置额外的头
func setHeaders(r *http.Request) {
	// 追踪头
	if cfg.Cfg.Proxy.TraceHeader != "" {
		r.Header.Set(cfg.Cfg.Proxy.TraceHeader, uuid.LowerV1WithoutHyphen())
	}
	// 地址头
	if cfg.Cfg.Proxy.IPAddrHeader != "" {
		r.Header.Set(cfg.Cfg.Proxy.IPAddrHeader, r.RemoteAddr)
	}
}

// ReLoadDB 重新加载数据直到成功
//
// todo 如果是高可用部署，在另外一个 apigateway 修改了数据
// 得想个办法让这边也重新加载一下数据
//
// 通过 redis/etcd watch key 的方式，但是不一定成功
func ReLoadDB() {
	go reloadDBRoutine()
}

// reload 重新加载数据
func reloadDB() error {
	// 查询数据库
	serviceModels, err := db.ServiceDA.All(&db.ServiceQuery{
		Enable: &db.True,
	})
	if err != nil {
		return err
	}
	serverModels, err := db.ServerDA.All(&db.ServerQuery{
		Enable: &db.True,
	})
	if err != nil {
		return err
	}
	// 更新内存
	ss := new(services)
	ss.init()
	for _, serviceModel := range serviceModels {
		_ss := ss.add(serviceModel.Path)
		for _, serverModel := range serverModels {
			// 服务组相同
			if serverModel.ServiceID != serviceModel.ID {
				continue
			}
			_ss.add(serverModel)
		}
	}
	_services = ss
	//
	return nil
}

// reloadRoutine 在协程中加载数据
func reloadDBRoutine() {
	defer func() {
		// 异常
		log.Recover(recover())
		// 结束
	}()
	for {
		// 加载
		err := reloadDB()
		if err == nil {
			return
		}
		log.Error(err)
		// 错误，休息
		time.Sleep(time.Second)
	}
}
