package login

import (
	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/login")
	//
	// g.GET("", get)
	g.POST("", post)
}
