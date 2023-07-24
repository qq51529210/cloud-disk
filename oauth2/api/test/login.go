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
<a href="{{.Code}}">oauth2-授权码模式</a>
</br>
<a href="{{.Token}}">oauth2-隐式模式</a>
</br>
<a href="{{.Password}}">oauth2-密码模式</a>
</br>
<a href="{{.Credentials}}">oauth2-客户端凭证模式</a>
</body>
</html>`)
}

type loginTP struct {
	Code        string
	Token       string
	Password    string
	Credentials string
}

func (tp *loginTP) initCode() {
	query := make(url.Values)
	query.Set("client_id", client)
	query.Set("scope", "avatar name friends")
	query.Set("state", state)
	query.Set("response_type", "code")
	query.Set("redirect_uri", fmt.Sprintf("http://%s/oauth2?response_type=code", cfg.Cfg.Test))
	tp.Code = fmt.Sprintf("http://%s/oauth2/authorize?%s", cfg.Cfg.Addr, query.Encode())
}

func (tp *loginTP) initToken() {
	query := make(url.Values)
	query.Set("client_id", client)
	query.Set("scope", "avatar name")
	query.Set("state", state)
	query.Set("response_type", "token")
	query.Set("redirect_uri", fmt.Sprintf("http://%s/oauth2?response_type=token", cfg.Cfg.Test))
	tp.Token = fmt.Sprintf("http://%s/oauth2/authorize?%s", cfg.Cfg.Addr, query.Encode())
}

func (tp *loginTP) initPassword() {
	query := make(url.Values)
	query.Set("grant_type", "password")
	query.Set("client_id", client)
	query.Set("client_secret", pwd)
	query.Set("username", user)
	query.Set("password", pwd)
	query.Set("scope", "avatar name")
	tp.Password = fmt.Sprintf("http://%s/oauth2/token?%s", cfg.Cfg.Addr, query.Encode())
}

func (tp *loginTP) initCredentials() {
	query := make(url.Values)
	query.Set("grant_type", "client_credentials")
	query.Set("client_id", client)
	query.Set("client_secret", pwd)
	query.Set("scope", "avatar name")
	tp.Credentials = fmt.Sprintf("http://%s/oauth2/token?%s", cfg.Cfg.Addr, query.Encode())
}

func login(ctx *gin.Context) {
	var t loginTP
	t.initCode()
	t.initToken()
	t.initPassword()
	t.initCredentials()
	tp.Execute(ctx.Writer, &t)
}
