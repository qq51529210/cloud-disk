package token

import (
	"auth/api/internal"
	"auth/db"

	"github.com/gin-gonic/gin"
)

type postReq struct {
	// 指定用于获取令牌的授权类型
	GrantType string `form:"grant_type" binding:"oneof=authorization_code client_credentials password"`
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

func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 查询
	app, err := db.GetApp(req.ClientID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if app == nil {
		internal.Data404(ctx)
		return
	}
	// 跳转
}
