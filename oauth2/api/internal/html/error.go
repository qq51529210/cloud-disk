package html

import (
	"html/template"
	"io"
)

var (
	error *template.Template
)

func init() {
	error, _ = template.New("error").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.Title}}</title>
</head>
<body>
	<h4>{{.Phrase}}</h4>
	<p>{{.Detail}}</p>
</body>
</html>
`)
}

// Error 用于格式化 error 模板
type Error struct {
	Title  string
	Phrase string
	Detail string
}

// Exec 格式化
func (m *Error) Exec(w io.Writer) {
	error.Execute(w, m)
}

// ExecError 用于格式化 error 模板
func ExecError(w io.Writer, title, phrase, detail string) {
	m := &Error{
		Title:  title,
		Phrase: phrase,
		Detail: detail,
	}
	m.Exec(w)
}
