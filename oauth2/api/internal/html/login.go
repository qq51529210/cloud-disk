package html

import (
	"html/template"
	"io"
)

var (
	login *template.Template
)

func init() {
	login, _ = template.New("login").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>登录</title>
	<style>` + css + `
	input[type="text"],
	input[type="password"] {
		width: 378px;
		padding: 10px;
		border: 1px solid #ccc;
		border-radius: 3px;
	}
	</style>
</head>
<body>
	<div class="container">
	<h2>登录</h2>
	<form method="post" action="{{.Action}}">
		<div class="form-group">
		<label for="username">用户名</label>
		<input type="text" name="account" value="test-user" required>
		</div>
		<div class="form-group">
		<label for="password">密码</label>
		<input type="password" name="password" value="123123" required>
		</div>
        <input type="hidden" name="redirect_uri" value="{{.RedirectURI}}">
		<button type="submit">确定</button>
	</form>
	</div>
</body>
</html>
`)
}

// Login 用于格式化 login 模板
type Login struct {
	Action      string
	RedirectURI string
}

// Exec 格式化
func (m *Login) Exec(w io.Writer) {
	login.Execute(w, m)
}

// ExecLogin 用于格式化 login 模板
func ExecLogin(w io.Writer, action, redirectURI string) {
	m := &Login{
		Action:      action,
		RedirectURI: redirectURI,
	}
	m.Exec(w)
}
