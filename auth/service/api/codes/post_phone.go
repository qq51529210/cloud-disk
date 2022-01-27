package codes

import (
	"encoding/json"
	"net/http"

	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/micro-services/auth/service"
	"github.com/qq51529210/micro-services/auth/util"
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
	code, err := cache.GetCache().NewPhoneCode(m1.Number)
	if err != nil {
		service.ParseJSONError(ctx, err)
		return
	}
	// 发送
	err = util.SenSMS(m1.Number, code)
	if err != nil {
		log.Error(err)
	}
	// 返回
	ctx.ResponseWriter.WriteHeader(http.StatusCreated)
	log.Infof("<%s> <%s>", m1.Number, code)
}
