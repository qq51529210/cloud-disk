package tokens

import (
	"github.com/qq51529210/micro-services/auth/service/middleware"
	"github.com/qq51529210/web/router"
)

func Init(r router.Router) {
	r.DELETE("?", delete)
	r.POST("", middleware.ParseForm, post)
}
