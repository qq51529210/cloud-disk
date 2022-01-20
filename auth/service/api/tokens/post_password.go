package tokens

import (
	"encoding/json"
	"net/http"

	"github.com/qq51529210/log"
	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/micro-services/auth/service"
	"github.com/qq51529210/micro-services/auth/store"
	"github.com/qq51529210/uuid"
	"github.com/qq51529210/web/router"
	"github.com/qq51529210/web/util"
)

type accountModel struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func postPassword(ctx *router.Context) {
	// 解析JSON
	var m1 accountModel
	err := json.NewDecoder(ctx.Request.Body).Decode(&m1)
	if err != nil {
		service.ParseJSONError(ctx, err)
		return
	}
	// 查询数据库
	m2, err := store.GetStore().UserStore().Get(m1.Account)
	if err != nil {
		service.QueryDataError(ctx, err)
		return
	}
	// 比较密码
	m1.Password = util.SHA1String(m1.Password)
	if m1.Password != m2.Password {
		service.UnauthorizedError(ctx)
		return
	}
	// 创建token
	token := uuid.LowerV1WithoutHyphen()
	err = cache.Set(token, m2, service.TokenExpire)
	if err != nil {
		service.QueryDataError(ctx, err)
		return
	}
	// 返回
	ctx.JSON(http.StatusCreated, map[string]string{
		"token": token,
	})
	log.Infof("postPassword: <%s> <%s> <%s>", m1.Account, m1.Password, token)
}
