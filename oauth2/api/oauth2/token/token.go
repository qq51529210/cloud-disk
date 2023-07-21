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
	case GrantTypeImplicit:
	case GrantTypePassword:
	case GrantTypeClientCredentials:
	case GrantTypeRefreshToken:
	}
}
