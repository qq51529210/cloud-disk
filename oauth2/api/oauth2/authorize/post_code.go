package authorize

import (
	"net/http"
	"net/url"
	"oauth2/api/internal/html"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

// postCode 处理用户确认授权后的 code 流程
func postCode(ctx *gin.Context, req *postReq) {
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session[*db.User])
	// 授权码
	code := new(db.AuthorizationCode)
	code.Scope = parsePostScope(ctx, req.form.Client)
	code.Client = req.form.Client
	code.UserID = sess.Data.ID
	util.CopyStruct(code, req)
	err := db.PutAuthorizationCode(code)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorDB, err.Error())
		return
	}
	// 重定向地址
	_u, err := url.Parse(req.RedirectURI)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorQuery, err.Error())
		return
	}
	q := _u.Query()
	q.Set(queryNameState, req.State)
	q.Set(queryNameCode, code.ID)
	_u.RawQuery = q.Encode()
	// 跳转
	ctx.Redirect(http.StatusSeeOther, _u.String())
}
