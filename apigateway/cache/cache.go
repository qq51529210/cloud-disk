package cache

import (
	"apigateway/db"
	"time"

	"github.com/qq51529210/log"
)

// Init 初始化
func Init() error {
	return reload()
}

// ReLoad 重新加载数据直到成功
// todo 如果是高可用部署，在另外一个 apigateway 修改了数据
// 得想个办法让这边也重新加载一下数据
//
// 通过 redis/etcd watch key 的方式，但是不一定成功
func ReLoad() {
	go reloadRoutine()
}

// reload 重新加载数据
func reload() error {
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
func reloadRoutine() {
	defer func() {
		// 异常
		log.Recover(recover())
		// 结束
	}()
	for {
		// 加载
		err := reload()
		if err == nil {
			return
		}
		log.Error(err)
		// 错误，休息
		time.Sleep(time.Second)
	}
}
