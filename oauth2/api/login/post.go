package login

import (
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

type postReq struct {
	// 账号
	Account string `form:"account" binding:"required,max=40"`
	// 密码
	Password string `form:"password" binding:"required,max=40"`
}

func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 数据库
	user, err := db.GetUserByAccount(req.Account)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if user == nil {
		internal.Data404(ctx)
		return
	}
	// 会话
	sess, err := db.NewSession(user)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	ctx.SetCookie(middleware.CookieName, sess.ID, -1, "/", "", true, true)
}
