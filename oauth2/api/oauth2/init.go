package oauth2

import (
	"oauth2/api/oauth2/authorize"
	"oauth2/api/oauth2/token"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/oauth2")
	//
	authorize.Init(g)
	token.Init(g)
}
