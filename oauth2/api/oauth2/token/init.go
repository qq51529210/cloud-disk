package token

import (
	"github.com/gin-gonic/gin"
)

const (
	
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/token")
	//
	g.GET("", token)
	g.POST("", token)
}
