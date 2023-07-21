package token

import (
	"oauth2/api/internal/html"

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

var (
	grantTypeHandle = make(map[string]func(*gin.Context))
)

func init() {
	grantTypeHandle[GrantTypeAuthorizationCode] = authorizationCode
	// grantTypeHandle[GrantTypePassword] = password
	// grantTypeHandle[GrantTypeClientCredentials] = clientCredentials
	// grantTypeHandle[GrantTypeImplicit] = implicit
	// grantTypeHandle[GrantTypeRefreshToken] = refreshToken
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

type tokenRes struct {
	// 应用程序在请求访问受保护资源时使用的令牌。
	// 它代表了客户端被授权的权限
	AccessToken string `json:"AccessToken"`
	// 该字段指示返回的令牌类型。
	// 比如 Bearer 令牌，意味着客户端可以简单地在后续请求的 "Authorization" 头中附上该令牌
	TokenType string `json:"token_type"`
	// 该字段以秒为单位指定访问令牌的过期时间。
	// 在此时间之后，访问令牌将不再有效，客户端需要获取新的访问令牌。
	Expires int64 `json:"expires_in"`
	// 该令牌可由客户端用于在当前访问令牌过期时获取新的访问令牌。
	// 通常在OAuth2的刷新令牌流程中使用，
	// 以便在不需要用户重新认证的情况下获取新的访问令牌
	RefreshToken string `json:"refresh_token"`
}

// token 处理获取访问令牌
func token(ctx *gin.Context) {
	// 处理
	hd, ok := grantTypeHandle[ctx.Query("grant_type")]
	if !ok {
		html.ExecError(ctx.Writer, html.TitleAccessToken, html.ErrorQuery, "")
		return
	}
	hd(ctx)
}

type authorizationCodeReq struct {
	// 在授权码模式中使用，表示从授权服务器获取的授权码
	Code string `form:"code" binding:"required"`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required,max=40"`
	// 表示客户端应用程序的密钥，由授权服务器分配给客户端
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
	// 重定向 URL
	RedirectURI string `form:"redirect_uri" binding:"uri"`
}

// authorizationCode 处理 grant_type=authorization_code
func authorizationCode(ctx *gin.Context) {
	// // 参数
	// var req authorizationCodeReq
	// err := ctx.ShouldBindQuery(&req)
	// if err != nil {
	// 	internal.Submit400(ctx, err.Error())
	// 	return
	// }
	// // 验证
	// ok, err := db.GetAuthorizationCode(req.Code)
	// if err != nil {
	// 	internal.DB500(ctx, err)
	// 	return
	// }
	// if !ok {
	// 	internal.Data404(ctx)
	// }
	// // 查询
	// Client, err := db.GetClient(req.ClientID)
	// if err != nil {
	// 	internal.DB500(ctx, err)
	// 	return
	// }
	// // 应用
	// if Client == nil || *Client.Enable != db.True || *Client.Secret != req.ClientSecret {
	// 	internal.Data404(ctx)
	// 	return
	// }
	// // 颁发
}
