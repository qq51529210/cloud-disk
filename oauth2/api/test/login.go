package test

import (
	"fmt"
	"html/template"
	"net/url"
	"oauth2/cfg"

	"github.com/gin-gonic/gin"
)

var (
	tp *template.Template
)

func init() {
	tp, _ = template.New("test").Parse(`<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>测试 oauth2 登录</title>
</head>
<body>
<a href="{{.Code}}">oauth2-授权码登录</a>
</br>
<a href="{{.Token}}">oauth2-隐式登录</a>
</body>
</html>`)
}

type loginTP struct {
	Code  string
	Token string
}

func login(ctx *gin.Context) {
	query := make(url.Values)
	query.Set("client_id", client)
	query.Set("scope", "avatar name friends")
	query.Set("state", state)
	query.Set("redirect_uri", fmt.Sprintf("%s/oauth2", cfg.Cfg.Test))
	//
	var t loginTP
	query.Set("response_type", "code")
	t.Code = fmt.Sprintf("http://%s/oauth2/authorize?%s", cfg.Cfg.Addr, query.Encode())
	query.Set("response_type", "token")
	t.Token = fmt.Sprintf("http://%s/oauth2/authorize?%s", cfg.Cfg.Addr, query.Encode())
	tp.Execute(ctx.Writer, &t)
}
