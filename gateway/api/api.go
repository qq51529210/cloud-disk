package api

import (
	"gateway/api/internal/middleware"
	"gateway/api/services"
	"gateway/cfg"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

var (
	// gin
	g = gin.New()
)

// Serve 开始服务
func Serve(staticsDir fs.FS) error {
	gin.SetMode(gin.DebugMode)
	// 静态文件
	err := util.GinStaticDir(g, "", "", staticsDir)
	if err != nil {
		return err
	}
	// 路由
	initRouter()
	// 监听
	return http.ListenAndServe(cfg.Cfg.AdminAddr, g)
}

// 初始化路由
func initRouter() {
	// 全局
	g.Use(middleware.Log)
	// 测试，充当代理的认证服务
	g.GET("/token", func(ctx *gin.Context) {})
	//
	services.Init(g)
}
