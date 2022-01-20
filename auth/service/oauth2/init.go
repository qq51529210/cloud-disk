package oauth2

import (
	"github.com/qq51529210/web/router"
)

func Init(r router.Router) {
	r.GET("authentication", parseAuthenticationQuery, getAuthentication)
	r.GET("authentication", postAuthentication)
	r.GET("authorization", getAuthorization)
}
