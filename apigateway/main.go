package main

import (
	"apigateway/api"
	"apigateway/cache"
	"apigateway/cfg"
	"apigateway/db"
	"embed"

	"github.com/qq51529210/log"
	"github.com/qq51529210/util"
)

//go:embed html/dist
var staticDir embed.FS

// @Title   接口文档
// @version 1.0.0
func main() {
	defer func() {
		log.Recover(recover())
	}()
	// 配置
	err := cfg.Load()
	if err != nil {
		panic(err)
	}
	// 日志
	err = util.InitLog(&cfg.Cfg.Log)
	if err != nil {
		panic(err)
	}
	// 数据库
	err = db.Init()
	if err != nil {
		panic(err)
	}
	// 缓存
	err = cache.Init()
	if err != nil {
		panic(err)
	}
	// 服务
	api.Serve(staticDir)
}
