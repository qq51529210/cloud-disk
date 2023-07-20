package authorize

import (
	"net/http"
	"net/url"
	"oauth2/api/internal/html"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

const (
	queryNameState = "state"
	queryNameCode  = "code"
)

type postReq struct {
	Name         string `form:"name"`
	Image        string `form:"image"`
	ResponseType string `form:"response_type" binding:"required,oneof=code token"`
	State        string `form:"state"`
	RedirectURI  string `form:"redirect_uri" binding:"uri"`
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
	// 授权码
	code, err := db.NewAuthorizationCodeTimeout()
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
	q.Set(queryNameCode, code)
	// 跳转
	_u.RawQuery = q.Encode()
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
