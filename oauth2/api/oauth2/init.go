package oauth2

import (
	"auth/api/oauth2/authorize"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/oauth2")
	//
	authorize.Init(g)
}
