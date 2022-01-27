package codes

import (
	"github.com/qq51529210/micro-services/auth/service/middleware"
	"github.com/qq51529210/web/router"
)

type Config struct {
	PhoneCodeExpire int
}

var (
	_Config Config
)

func Init(r router.Router, cfg *Config) {
	_Config = *cfg
	if _Config.PhoneCodeExpire < 1 {
		_Config.PhoneCodeExpire = 60
	}
	r.POST("", middleware.ParseForm, post)
}
