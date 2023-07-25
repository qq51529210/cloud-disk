package proxy

import (
	"apigateway/cache"
	"apigateway/cfg"
	"net/http"
	"time"

	"github.com/qq51529210/log"
	"github.com/qq51529210/uuid"
)

// Serve 开始服务
func Serve() error {
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
	// 身份验证
	if ser.Auth && !authorization(w, r) {
		return
	}
	// 设置头
	setHeaders(r)
	// 转发
	ser.ServeHTTP(w, r)
}

// authorization 身份验证
func authorization(w http.ResponseWriter, r *http.Request) bool {
	return true
}

// getServer 返回代理的服务
func getServer(r *http.Request) *cache.Server {
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
	return cache.GetMinLoadServer(path)
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
