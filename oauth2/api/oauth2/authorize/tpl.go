package authorize

import "html/template"

var (
	errorTP     *template.Template
	authorizeTP *template.Template
)

func init() {
	errorTP, _ = template.New("error").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>授权</title>
</head>
<body>
	<p>{{.}}</p>
</body>
</html>
`)

	authorizeTP, _ = template.New("authorize").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>授权</title>
	<style>
	body {
		font-family: Arial, sans-serif;
		background-color: #f4f4f4;
		padding: 20px;
	}
	
	.container {
		max-width: 400px;
		margin: 0 auto;
		background-color: #ffffff;
		padding: 20px;
		border-radius: 5px;
		box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
	}
	
	.form-group {
		margin-bottom: 20px;
	}
	
	label {
		display: block;
		font-weight: bold;
		margin-bottom: 5px;
	}
	
	button {
		display: block;
		width: 100%;
		padding: 10px;
		background-color: #4caf50;
		border: none;
		color: #ffffff;
		font-size: 16px;
		font-weight: bold;
		border-radius: 3px;
		cursor: pointer;
	}
	</style>
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
