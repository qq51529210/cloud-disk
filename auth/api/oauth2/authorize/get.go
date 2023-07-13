package authorize

import (
	"github.com/gin-gonic/gin"
)

// getReq 表示 authorize 提交的参数
type getReq struct {
	// 指定用于授权流程的响应类型，常见的值包括
	// code ，用于授权码授权流程
	// token ，用于隐式授权流程
	ResponseType string `form:"response_type" binding:"oneof=code token"`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required"`
	// 指定授权服务器将用户重定向到的客户端应用程序的回调 URL
	RedirectURI string `form:"redirect_uri" binding:"required,uri"`
	// 表示客户端请求的访问权限范围
	Scope string `form:"scope"`
	// 用于在授权请求和授权响应之间传递状态，以防止 CSRF 攻击
	State string `form:"state"`
	//	grant_type：用于请求访问令牌的授权类型，常见的值包括authorization_code（授权码授权流程）、password（密码授权流程）和refresh_token（刷新令牌）。
	//
	// 7. code：用于在授权码授权流程中交换访问令牌的授权码。
	// 8. access_token：表示访问令牌，用于访问受保护的资源。
	// 9. refresh_token：表示刷新令牌，用于获取新的访问令牌
}

func get(ctx *gin.Context) {
	// 返回授权确认页面
}
