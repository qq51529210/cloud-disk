package oauth2

import (
	"github.com/qq51529210/web/router"
)

func Init(r router.Router) {
	r.POST("/code")
}
