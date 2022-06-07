package main

import (
	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-service/authentication/api"
	"github.com/qq51529210/micro-service/authentication/config"
	"github.com/qq51529210/web"
)

func main() {
	defer func() {
		log.Recover(recover())
	}()
	log.Info("", "app start")
	// 启动配置参数
	cfg := config.Load()
	// 路由
	handler := api.New()
	// 初始化服务
	var ser web.Server
	if cfg.HTTP.CertPEM != "" && cfg.HTTP.KeyPEM != "" {
		ser = web.NewTLSServerWithKeyPair(cfg.HTTP.Address, []byte(cfg.HTTP.CertPEM), []byte(cfg.HTTP.KeyPEM), handler)
	} else {
		ser = web.NewServer(cfg.HTTP.Address, handler)
	}
	// 监听
	ser.Serve()
}
