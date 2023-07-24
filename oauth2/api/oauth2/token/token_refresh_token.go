package token

import (
	"oauth2/api/internal"
	"oauth2/api/internal/html"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/log"
)

type tokenRefreshTokenReq struct {
	ClientID     string `form:"client_id" binding:"required,max=40"`
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
	RefreshToken string `form:"refresh_token" binding:"required,max=40"`
}

// tokenRefreshToken 处理 grant_type=password
func tokenRefreshToken(ctx *gin.Context) {
	// 参数
	var req tokenRefreshTokenReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 刷新令牌
	refreshToken, err := db.GetRefreshToken(req.RefreshToken)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if refreshToken == nil {
		internal.Submit400(ctx, "无效刷新令牌")
		return
	}
	// 应用
	client, err := db.GetClient(req.ClientID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if client == nil || client.ID != refreshToken.ClientID || *client.Enable != db.True {
		internal.Submit400(ctx, html.ErrorClientNotFound)
		return
	}
	if *client.Secret != req.ClientSecret {
		internal.Submit400(ctx, html.ErrorClientSecret)
		return
	}
	// 访问令牌
	accessToken := new(db.Token)
	accessToken.TokenType = *client.TokenType
	accessToken.Scope = refreshToken.Scope
	accessToken.GrantType = db.GrantTypeRefreshToken
	accessToken.UserID = refreshToken.UserID
	accessToken.ClientID = client.ID
	err = db.PutAccessToken(accessToken)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	onOK(ctx, accessToken)
	// 删除旧的刷新令牌
	err = db.DelRefreshToken(refreshToken.AccessToken)
	if err != nil {
		log.Errorf("del refresh token %s error: %s", refreshToken.AccessToken, err.Error())
	}
}
