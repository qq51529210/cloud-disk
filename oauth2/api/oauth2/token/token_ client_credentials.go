package token

import (
	"oauth2/api/internal"
	"oauth2/api/internal/html"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

type tokenClientCredentialsReq struct {
	ClientID     string `form:"client_id" binding:"required,max=40"`
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
	Scope        string `form:"scope" binding:"required"`
}

// tokenPassword 处理 grant_type=client_credentials
func tokenClientCredentials(ctx *gin.Context) {
	// 参数
	var req tokenPasswordReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 应用
	client, err := db.GetClient(req.ClientID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if client == nil || *client.Enable != db.True {
		internal.Submit400(ctx, html.ErrorClientNotFound)
		return
	}
	// 令牌
	token := new(db.AccessToken)
	token.Type = *client.TokenType
	token.Scope = req.Scope
	token.GenType = db.GenTypeCredentials
	token.ClientID = client.ID
	err = db.PutAccessToken(token)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	onOK(ctx, token)
}
