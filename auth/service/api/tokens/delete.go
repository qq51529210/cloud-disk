package tokens

import (
	"net/http"

	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/micro-services/auth/service"
	"github.com/qq51529210/web/router"
)

func delete(ctx *router.Context) {
	err := cache.Del(ctx.Param[0])
	if err != nil {
		service.QueryDataError(ctx, err)
		return
	}
	ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	log.Infof("delete token <%s>", ctx.Param[0])
}
