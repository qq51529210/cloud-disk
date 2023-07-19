package login

import (
	"net/http"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
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
	if user == nil || *user.Enable != 1 {
		internal.Data404(ctx)
		return
	}
	if *user.Password != util.SHA1String(req.Password) {
		ctx.JSON(http.StatusBadRequest, &internal.Error{
			Phrase: "登录失败",
			Detail: "账号或密码不正确",
		})
		return
	}
	// 会话
	sess, err := db.NewSession(user)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// cookie
	ctx.SetCookie(middleware.CookieName, sess.ID, -1, "/", "", true, true)
	// 返回
}
