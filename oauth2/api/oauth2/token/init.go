package token

import (
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/token", middleware.CheckSession)
	//
	g.POST("", post)
}
