package html

import (
	"html/template"
	"io"
)

var (
	authorize *template.Template
)

func init() {
	authorize, _ = template.New("authorize").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>授权</title>
	<style>` + css + `</style>
</head>
<body>
	<div class="container">
	<h2>访问授权</h2>
	<p>应用<h4>[{{.ClientName}}]</h4>
	<img src="{{.ClientImage}}" width="16" height="16">
	请求访问以下的数据</p>
	<form method="post" action="/oauth2/authorize">
	{{range .Scope}}
		<label>
			<input type="checkbox" name="{{.Key}}" {{if .Check}}checked{{end}}> {{.Name}}
		</label>
	{{end}}
		<input type="hidden" name="response_type" value="{{.ResponseType}}">
		<input type="hidden" name="client_id" value="{{.ClientID}}">
		<input type="hidden" name="state" value="{{.State}}">
		<input type="hidden" name="redirect_uri" value="{{.RedirectURI}}">
		<button type="submit">确定</button>
	</form>
	</div>
</body>
</html>
`)
}

// AuthorizeScope 表示 Authorize 的 Scope 字段
type AuthorizeScope struct {
	// name="key"
	Key string
	// label
	Name string
	// 是否选中
	Check bool
}

// Authorize 用于格式化 authorize 模板
type Authorize struct {
	ClientName   string
	ClientImage  string
	ResponseType string
	ClientID     string
	State        string
	RedirectURI  string
	Scope        []*AuthorizeScope
}

// Exec 格式化
func (m *Authorize) Exec(w io.Writer) {
	authorize.Execute(w, m)
}
