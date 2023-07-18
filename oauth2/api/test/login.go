package test

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	query := make(url.Values)
	query.Set("response_type", "code")
	query.Set("client_id", app)
	query.Set("scope", "readwrite")
	query.Set("state", state)
	query.Set("redirect_uri", fmt.Sprintf("%s/oauth2", host))
	redirectURL := fmt.Sprintf("%s/authorize?%s", oauth2Host, query.Encode())
	ctx.Redirect(http.StatusFound, redirectURL)
}
