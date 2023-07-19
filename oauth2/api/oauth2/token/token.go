package token

import (
	"net/http"
	"oauth2/api/internal"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

type tokenReq struct {
	// 指定用于获取令牌的授权类型
	GrantType string `form:"grant_type" binding:"oneof=authorization_code client_credentials password refresh_token implicit"`
	// 在授权码模式中使用，表示从授权服务器获取的授权码
	Code string `form:"code" binding:""`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required,max=40"`
	// 表示客户端应用程序的密钥，由授权服务器分配给客户端
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
	// // 用户名，密码模式中使用，表示资源所有者的用户名和密码，用于直接获取令牌
	// Username string `form:"username" binding:"required,max=40"`
	// // 密码，密码模式中使用，表示资源所有者的用户名和密码，用于直接获取令牌
	// Password string `form:"password" binding:""`
}

type tokenRes struct {
	// 应用程序在请求访问受保护资源时使用的令牌。
	// 它代表了客户端被授权的权限
	AccessToken string `json:"AccessToken"`
	// 该字段指示返回的令牌类型。
	// 比如 Bearer 令牌，意味着客户端可以简单地在后续请求的 "oauth2orization" 头中附上该令牌
	TokenType string `json:"token_type"`
	// 该字段以秒为单位指定访问令牌的过期时间。
	// 在此时间之后，访问令牌将不再有效，客户端需要获取新的访问令牌。
	Expires int64 `json:"expires_in"`
	// 该令牌可由客户端用于在当前访问令牌过期时获取新的访问令牌。
	// 通常在OAuth2的刷新令牌流程中使用，
	// 以便在不需要用户重新认证的情况下获取新的访问令牌
	RefreshToken string `json:"refresh_token"`
}

func token(ctx *gin.Context) {
	// 参数
	var req tokenReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 查询
	Client, err := db.GetClient(req.ClientID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if Client == nil || *Client.Enable != db.True || *Client.Secret != req.ClientSecret {
		internal.Data404(ctx)
		return
	}
	// 返回
	var res tokenRes
	ctx.JSON(http.StatusOK, &res)
}
