package login

import (
	"html/template"
	"net/http"
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
		postError(ctx, "参数错误")
		return
	}
	// 数据库
	user, err := db.GetUserByAccount(req.Account)
	if err != nil {
		postError(ctx, "数据库错误")
		return
	}
	if user == nil || *user.Enable != 1 ||
		*user.Password != util.SHA1String(req.Password) {
		postError(ctx, "账号或密码不正确")
		return
	}
	// 会话
	sess, err := db.NewUserSession(user)
	if err != nil {
		postError(ctx, "数据库错误")
		return
	}
	// cookie
	ctx.SetCookie(middleware.CookieName, sess.ID, int(sess.Expires), "/", "", true, true)
	// 跳转
	redirectURL := ctx.Query(middleware.QueryRedirectURI)
	if redirectURL != "" {
		ctx.Redirect(http.StatusSeeOther, redirectURL)
	} else {
		ctx.Redirect(http.StatusSeeOther, "/")
	}
}

var (
	loginErrorTP *template.Template
)

func init() {
	loginErrorTP, _ = template.New("loginError").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>登录</title>
</head>
<body>
	<p>{{.}}</p>
	<a href="javascript:history.back()">返回登录</a>
</body>
</html>
`)
}

func postError(ctx *gin.Context, err string) {
	loginErrorTP.Execute(ctx.Writer, err)
}
