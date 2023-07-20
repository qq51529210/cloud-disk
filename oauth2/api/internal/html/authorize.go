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
	<p>应用<h4>[{{.ClientName}}]</h4>请求访问以下的数据</p>
	<form method="post" action="{{.Action}}">
	{{range $key, $value := .Scope}}
		<label>
			<input type="checkbox" name="{{$key}}"> {{$value}}
		</label>
	{{end}}
		<button type="submit">确定</button>
	</form>
	</div>
</body>
</html>
`)
}

// Authorize 用于格式化 authorize 模板
type Authorize struct {
	Action string
	Scope  map[string]string
}

// Exec 格式化
func (m *Authorize) Exec(w io.Writer) {
	error.Execute(w, m)
}
