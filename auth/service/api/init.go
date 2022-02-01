package api

import (
	"github.com/qq51529210/micro-services/auth/service/api/apps"
	"github.com/qq51529210/micro-services/auth/service/api/tokens"
	"github.com/qq51529210/micro-services/auth/service/api/users"
	"github.com/qq51529210/web/router"
)

func Init(r router.Router) {
	apps.Init(r.SubRouter("apps"))
	tokens.Init(r.SubRouter("tokens"))
	users.Init(r.SubRouter("users"))
}
