package token

import (
	"oauth2/api/internal"
	"oauth2/api/internal/html"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

type tokenPasswordReq struct {
	ClientID     string `form:"client_id" binding:"required,max=40"`
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
	Username     string `form:"username" binding:"required,max=40"`
	Password     string `form:"password" binding:"required,max=40"`
	Scope        string `form:"scope" binding:"required"`
}

// tokenPassword 处理 grant_type=password
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
	if *user.Password != util.SHA1String(req.Password) {
		internal.Submit400(ctx, html.ErrorUsernameOrPassword)
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
	token := new(db.Token)
	token.TokenType = *client.TokenType
	token.Scope = req.Scope
	token.GrantType = db.GenTypePassword
	token.UserID = user.ID
	token.ClientID = client.ID
	err = db.PutAccessToken(token)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	onOK(ctx, token)
}
