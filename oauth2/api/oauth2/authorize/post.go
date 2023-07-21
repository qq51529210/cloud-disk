package authorize

import (
	"net/http"
	"net/url"
	"oauth2/api/internal/html"
	"oauth2/api/internal/middleware"
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

const (
	queryNameState = "state"
	queryNameCode  = "code"
)

type postReq struct {
	Name         string `form:"name"`
	Image        string `form:"image"`
	ResponseType string `form:"response_type" binding:"required,oneof=code token"`
	ClientID     string `form:"client_id" binding:"required,max=40"`
	State        string `form:"state"`
	RedirectURI  string `form:"redirect_uri" binding:"uri"`
}

func parsePostScope(ctx *gin.Context) string {
	var scope strings.Builder
	for k := range ctx.Request.PostForm {
		switch k {
		case "response_type", "client_id", "state", "redirect_uri":
		default:
			scope.WriteString(k)
			scope.WriteByte(' ')
		}
	}
	return scope.String()
}

func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorQuery, err.Error())
		return
	}
	// 模式
	switch req.ResponseType {
	case ResponseTypeCode:
		postCode(ctx, &req)
	case ResponseTypeToken:
		postToken(ctx, &req)
	}
}

// postCode 处理用户确认授权后的 code 流程
func postCode(ctx *gin.Context, req *postReq) {
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session[*db.User])
	// 授权码
	code := new(db.AuthorizationCode)
	code.Scope = parsePostScope(ctx)
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

// postToken 处理用户确认授权后的 token 流程
func postToken(ctx *gin.Context, req *postReq) {
	//
	// // 跳转
	// redirectURL := ctx.Query(middleware.QueryRedirectURI)
	// if redirectURL != "" {
	// 	_u, err := url.Parse(redirectURL)
	// 	if err != nil {
	// 		errorTP.Execute(ctx.Writer, "第三方应用数据错误，无法完成跳转")
	// 		return
	// 	}
	// 	q := _u.Query()
	// 	q.Set(stateQueryName, ctx.Query(stateQueryName))
	// 	q.Set(codeQueryName, uuid.SnowflakeIDString())
	// 	_u.RawQuery = q.Encode()
	// 	ctx.Redirect(http.StatusSeeOther, _u.String())
	// }
}
