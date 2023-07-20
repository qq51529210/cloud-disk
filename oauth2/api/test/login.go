package test

import (
	"fmt"
	"html/template"
	"net/url"

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
<a href="{{.}}">oauth2登录</a>
</body>
</html>`)
}

func login(ctx *gin.Context) {
	responseType := ctx.Query("response_type")
	if responseType == "" {
		responseType = "code"
	}
	query := make(url.Values)
	query.Set("response_type", responseType)
	query.Set("client_id", client)
	query.Set("scope", "image name")
	query.Set("state", state)
	query.Set("redirect_uri", fmt.Sprintf("%s/oauth2", host))
	redirectURL := fmt.Sprintf("%s/oauth2/authorize?%s", oauth2Host, query.Encode())
	tp.Execute(ctx.Writer, redirectURL)
}
