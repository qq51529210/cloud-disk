package middleware

import (
	"github.com/qq51529210/log"
	"github.com/qq51529210/web/router"
)

func ParseForm(ctx *router.Context) {
	err := ctx.ParseForm()
	if err != nil {
		log.Error(err)
		ctx.Abort()
		return
	}
}
