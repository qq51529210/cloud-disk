package token

import (
	"oauth2/api/internal"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/log"
)

type tokenAuthorizationCodeReq struct {
	// 在授权码模式中使用，表示从授权服务器获取的授权码
	Code string `form:"code" binding:"required"`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required,max=40"`
	// 表示客户端应用程序的密钥，由授权服务器分配给客户端
	ClientSecret string `form:"client_secret" binding:"required,max=40"`
}

// tokenAuthorizationCode 处理 grant_type=authorization_code
func tokenAuthorizationCode(ctx *gin.Context) {
	// 参数
	var req tokenAuthorizationCodeReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 授权码
	code, err := db.GetAuthorizationCode(req.Code)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if code == nil || req.ClientID != code.Client.ID {
		internal.Submit400(ctx, "授权码错误")
		return
	}
	if *code.Client.Secret != req.ClientSecret {
		internal.Submit400(ctx, "应用密钥不正确")
		return
	}
	// 令牌
	token := new(db.AccessToken)
	token.Type = *code.Client.TokenType
	token.Scope = code.Scope
	token.GenType = db.GenTypeCode
	token.ClientID = code.Client.ID
	token.UserID = code.UserID
	err = db.PutAccessToken(token)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	onOK(ctx, token)
	// 删除授权码
	err = db.DelAuthorizationCode(code.ID)
	if err != nil {
		log.Errorf("del authorization code %s error: %s", code.ID, err.Error())
	}
}
