package codes

import (
	"encoding/json"
	"net/http"

	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/micro-services/auth/service"
	"github.com/qq51529210/uuid"
	"github.com/qq51529210/web/router"
)

type phoneModel struct {
	Number string `json:"number"`
}

func postPhone(ctx *router.Context) {
	// 解析JSON
	var m1 phoneModel
	err := json.NewDecoder(ctx.Request.Body).Decode(&m1)
	if err != nil {
		service.ParseJSONError(ctx, err)
		return
	}
	// 生成验证码
	code := uuid.LowerV1WithoutHyphen()
	err = cache.Set(m1.Number, code, int(_Config.PhoneCodeExpire))
	if err != nil {
		service.ParseJSONError(ctx, err)
		return
	}
	// 返回
	ctx.ResponseWriter.WriteHeader(http.StatusCreated)
	log.Infof("codes.postPhone: <%s> <%s>", m1.Number, code)
}
