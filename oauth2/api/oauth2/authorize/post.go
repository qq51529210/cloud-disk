package authorize

import (
	"net/http"
	"net/url"
	"oauth2/api/internal"
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
	baseQuery
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
	str := scope.String()
	if str != "" {
		// 去掉后面的空格
		str = str[:len(str)-1]
	}
	return str
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
	// 会话
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
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session[*db.User])
	// 应用
	client, err := db.GetClient(req.ClientID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if client == nil || *client.Enable != db.True {
		internal.Submit400(ctx, "应用不存在")
		return
	}
	// 令牌
	token := new(db.AccessToken)
	token.Type = *client.TokenType
	token.Scope = req.Scope
	token.ClientID = req.ClientID
	token.UserID = sess.Data.ID
	err = db.PutAccessToken(token)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 重定向
	if req.RedirectURI != "" {
		// 重定向地址
		_u, err := url.Parse(req.RedirectURI)
		if err != nil {
			internal.Submit400(ctx, err.Error())
			return
		}
		_u.RawQuery = util.HTTPQuery(token, _u.Query()).Encode()
		// 跳转
		ctx.Redirect(http.StatusSeeOther, _u.String())
		return
	}
	// 没有重定向，返回 JSON
	ctx.JSON(http.StatusOK, token)
}
