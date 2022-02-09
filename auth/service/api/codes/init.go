package codes

import (
	"github.com/qq51529210/micro-services/auth/service/middleware"
	"github.com/qq51529210/web/router"
)

func Init(r router.Router) {
	r.POST("", middleware.ParseForm, post)
}
