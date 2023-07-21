package login

import (
	"net/http"
	"oauth2/api/internal/html"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

// 登录方式
const (
	LoginTypePhone = "phone"
	LoginTypeEmail = "email"
)

// post 处理登录表单提交
func post(ctx *gin.Context) {
	switch ctx.Query("type") {
	case "":
		postAccount(ctx)
		// case LoginTypePhone:
		// case LoginTypeEmail:
	}
}

type postAccountReq struct {
	// 账号
	Account string `form:"account" binding:"required,max=40,alphanum"`
	// 密码
	Password string `form:"password" binding:"required"`
	// 重定向
	RedirectURI string `form:"redirect_uri" binding:"uri"`
}

// postAccount 处理默认
func postAccount(ctx *gin.Context) {
	// 参数
	var req postAccountReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleLogin, html.ErrorQuery, err.Error())
		return
	}
	// 数据库
	user, err := db.GetUserByAccount(req.Account)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleLogin, html.ErrorDB, err.Error())
		return
	}
	if user == nil ||
		*user.Enable != db.True ||
		*user.Password != util.SHA1String(req.Password) {
		html.ExecError(ctx.Writer, html.TitleLogin, "账号或密码不正确", "")
		return
	}
	// 会话
	sess, err := db.NewUserSession(user)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleLogin, html.ErrorDB, err.Error())
		return
	}
	// cookie
	ctx.SetCookie(middleware.CookieName, sess.ID, int(sess.Expires), "/", "", true, true)
	// 跳转
	if req.RedirectURI != "" {
		ctx.Redirect(http.StatusSeeOther, req.RedirectURI)
	} else {
		ctx.Redirect(http.StatusSeeOther, "/")
	}
}
