package login

import (
	"oauth2/api/internal/html"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	url := Path
	if ctx.Request.URL.RawQuery != "" {
		url += "?" + ctx.Request.URL.RawQuery
	}
	html.ExecLogin(ctx.Writer, url)
}
