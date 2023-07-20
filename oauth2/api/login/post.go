package login

import (
	"oauth2/api/internal/html"

	"github.com/gin-gonic/gin"
)

// 登录方式
const (
	LoginTypePhone = "phone"
	LoginTypeEmail = "email"
)

var (
	loginTypeHandle = make(map[string]func(*gin.Context))
)

func init() {
	loginTypeHandle[""] = account
}

func post(ctx *gin.Context) {
	// 处理
	hd, ok := loginTypeHandle[ctx.Query("type")]
	if !ok {
		html.ExecError(ctx.Writer, html.TitleLogin, html.ErrorQuery, "")
		return
	}
	hd(ctx)
}
