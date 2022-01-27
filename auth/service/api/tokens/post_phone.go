package tokens

import (
	"encoding/json"

	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/micro-services/auth/service"
	"github.com/qq51529210/micro-services/auth/store"
	"github.com/qq51529210/web/router"
)

type phoneModel struct {
	Number string `json:"number"`
	Code   string `json:"code"`
}

func postPhone(ctx *router.Context) {
	// 解析JSON
	var m1 phoneModel
	err := json.NewDecoder(ctx.Request.Body).Decode(&m1)
	if err != nil {
		service.ParseJSONError(ctx, err)
		return
	}
	// 检查验证码
	code, err := cache.GetPhoneCode().Get(m1.Number)
	if err != nil {
		service.ParseJSONError(ctx, err)
		return
	}
	if code != m1.Code {
		service.UnauthorizedError(ctx)
		return
	}
	// 查询数据库
	m2, err := store.GetUser().Get(m1.Number)
	if err != nil {
		service.QueryDataError(ctx, err)
		return
	}
	// 创建token
	token := createToken(ctx, m2)
	if token != "" {
		log.Infof("<%s> <%s> <%s>", m1.Number, m1.Code, token)
	}
}
