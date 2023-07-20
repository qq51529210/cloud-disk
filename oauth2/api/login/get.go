package login

import (
	"oauth2/api/internal/html"
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	html.ExecLogin(ctx.Writer, ctx.Query(middleware.QueryRedirectURI))
}
