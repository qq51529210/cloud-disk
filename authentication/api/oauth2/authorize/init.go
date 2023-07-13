package authorize

import (
	"authentication/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/authorize", middleware.CheckSession)
	//
	g.GET("", get)
	g.POST("", post)
}

//	grant_type：用于请求访问令牌的授权类型，常见的值包括authorization_code（授权码授权流程）、password（密码授权流程）和refresh_token（刷新令牌）。
//
// 7. code：用于在授权码授权流程中交换访问令牌的授权码。
// 8. access_token：表示访问令牌，用于访问受保护的资源。
// 9. refresh_token：表示刷新令牌，用于获取新的访问令牌
