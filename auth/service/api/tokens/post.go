package tokens

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/micro-services/auth/service"
	"github.com/qq51529210/micro-services/auth/store"
	"github.com/qq51529210/web/router"
)

func post(ctx *router.Context) {
	_type := ctx.Request.FormValue("type")
	switch _type {
	case "password", "":
		postPassword(ctx)
	case "phone":
		postPhone(ctx)
	default:
		service.FormValueError(ctx, "type", _type)
	}
}

func createToken(ctx *router.Context, model *store.UserModel) string {
	// 创建token
	var str strings.Builder
	json.NewEncoder(&str).Encode(model)
	token, err := cache.GetToken().New(str.String())
	if err != nil {
		service.QueryDataError(ctx, err)
		return ""
	}
	// 返回
	ctx.JSON(http.StatusCreated, map[string]string{
		"token": token,
	})
	return token
}
