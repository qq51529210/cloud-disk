package oauth2

import (
	"fmt"
	"net/http"

	"github.com/qq51529210/web/router"
)

type authenticationQuery struct {
	appId        string
	redirectUri  string
	responseType string
	scope        string
	state        string
}

func parseAuthenticationQuery(ctx *router.Context) {
	var query authenticationQuery
	query.appId = ctx.FormValue("app_id")
	if query.appId == "" {
		ctx.JSONBytes(http.StatusBadRequest, queryRequiredAppId)
		ctx.Abort()
		return
	}
	query.redirectUri = ctx.FormValue("redirect_uri")
	if query.redirectUri == "" {
		ctx.JSONBytes(http.StatusBadRequest, queryRequiredRedirectUri)
		ctx.Abort()
		return
	}
	query.responseType = ctx.FormValue("response_type")
	if query.responseType == "" {
		ctx.JSONBytes(http.StatusBadRequest, queryRequiredResponseType)
		ctx.Abort()
		return
	}
	if query.responseType != "token" {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error":   fmt.Sprintf("query <response_type> invalid value <%s>", query.responseType),
			"example": "response_type=token",
		})
		ctx.Abort()
		return
	}
	query.scope = ctx.FormValue("scope")
	if query.scope == "" {
		ctx.JSONBytes(http.StatusBadRequest, queryRequiredScope)
		ctx.Abort()
		return
	}
	query.state = ctx.FormValue("state")
	if query.state == "" {
		ctx.JSONBytes(http.StatusBadRequest, queryRequiredState)
		ctx.Abort()
		return
	}
	ctx.TempData = &query
}

// /authentication?app_id=xx&redirect_uri=xx&response_type=token&scope=xx&state=xx
func getAuthentication(ctx *router.Context) {
	query := ctx.TempData.(*authenticationQuery)
}
