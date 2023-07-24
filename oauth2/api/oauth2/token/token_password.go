package token

import (
	"oauth2/api/internal"
	"oauth2/api/internal/html"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

type tokenPasswordReq struct {
	ClientID     string `form:"client_id" binding:"required,max=40"`
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
	Username     string `form:"username" binding:"required,max=40"`
	Password     string `form:"password" binding:"required,max=40"`
	Scope        string `form:"scope" binding:"required"`
}

// tokenPassword 处理 grant_type=authorization_code
func tokenPassword(ctx *gin.Context) {
	// 参数
	var req tokenPasswordReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 用户
	user, err := db.GetUserByAccount(req.Username)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if user == nil || *user.Enable != db.True {
		internal.Submit400(ctx, html.ErrorUserNotFound)
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
	token.ClientID = client.ID
	token.UserID = user.ID
	err = db.PutAccessToken(token)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 继续处理
	onOK(ctx, token)
}
