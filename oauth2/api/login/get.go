package login

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

var (
	loginTP *template.Template
)

func init() {
	loginTP, _ = template.New("login").Parse(`
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>登录</title>
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
	
	input[type="text"],
	input[type="password"] {
		width: 378px;
		padding: 10px;
		border: 1px solid #ccc;
		border-radius: 3px;
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
	<h2>登录</h2>
	<form method="post" action="{{.}}">
		<div class="form-group">
		<label for="username">用户名</label>
		<input type="text" name="account" value="test-user" required>
		</div>
		<div class="form-group">
		<label for="password">密码</label>
		<input type="password" name="password" value="123123" required>
		</div>
		<button type="submit">确定</button>
	</form>
	</div>
</body>
</html>
`)
}

func get(ctx *gin.Context) {
	loginTP.Execute(ctx.Writer, fmt.Sprintf("%s?%s", Path, ctx.Request.URL.RawQuery))
}
