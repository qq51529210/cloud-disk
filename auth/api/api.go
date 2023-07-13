package api

import (
	"auth/api/apps"
	"auth/api/users"
	"auth/cfg"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// gin
	g *gin.Engine
)

// Serve 开始服务
func Serve() error {
	gin.SetMode(gin.ReleaseMode)
	g = gin.New()
	// 路由
	initRouter()
	// 监听
	return http.ListenAndServe(cfg.Cfg.Addr, g)
}

// 初始化路由
func initRouter() {
	apps.Init(g)
	users.Init(g)
}
