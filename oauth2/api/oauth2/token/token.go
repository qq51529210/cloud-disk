package token

import (
	"github.com/gin-gonic/gin"
)

// 模式
const (
	GrantTypeAuthorizationCode = "authorization_code"
	GrantTypePassword          = "password"
	GrantTypeClientCredentials = "client_credentials"
	GrantTypeImplicit          = "implicit"
	GrantTypeRefreshToken      = "refresh_token"
)

// token 处理获取访问令牌
func token(ctx *gin.Context) {
	switch ctx.Query("grant_type") {
	case GrantTypeAuthorizationCode:
		authorizationCode(ctx)
	case GrantTypePassword:
	case GrantTypeClientCredentials:
	case GrantTypeImplicit:
	case GrantTypeRefreshToken:
	}
}

// type tokenReq struct {
// 	// 指定用于获取令牌的授权类型
// 	GrantType string `form:"grant_type" binding:"oneof=authorization_code password client_credentials implicit refresh_token"`
// 	// 在授权码模式中使用，表示从授权服务器获取的授权码
// 	Code string `form:"code" binding:""`
// // 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
// ClientID string `form:"client_id" binding:"required,max=40"`
// // 表示客户端应用程序的密钥，由授权服务器分配给客户端
// ClientSecret string `form:"client_secret" binding:"required,max=40"`
// // 用户名，密码模式中使用，表示资源所有者的用户名和密码，用于直接获取令牌
// Username string `form:"username" binding:""`
// // 密码，密码模式中使用，表示资源所有者的用户名和密码，用于直接获取令牌
// Password string `form:"password" binding:""`
// }
