package authorize

import (
	"net/http"
	"net/url"
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/uuid"
)

const (
	stateQueryName = "state"
	codeQueryName  = "code"
)

func post(ctx *gin.Context) {
	// 跳转
	redirectURL := ctx.Query(middleware.QueryRedirectURI)
	if redirectURL != "" {
		_u, err := url.Parse(redirectURL)
		if err != nil {
			errorTP.Execute(ctx.Writer, "第三方应用数据错误，无法完成跳转")
			return
		}
		q := _u.Query()
		q.Set(stateQueryName, ctx.Query(stateQueryName))
		q.Set(codeQueryName, uuid.SnowflakeIDString())
		_u.RawQuery = q.Encode()
		ctx.Redirect(http.StatusSeeOther, _u.String())
	}
}
