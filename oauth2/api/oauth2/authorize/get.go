package authorize

import (
	"fmt"
	"oauth2/api/internal/html"
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

// 模式
const (
	ResponseTypeCode  = "code"
	ResponseTypeToken = "token"
	// ResponseTypePassword          = "password"
	// ResponseTypeClientCredentials = "client_credentials"
)

type getReq struct {
	// 指定用于授权流程的响应类型，常见的值包括
	// code 用于授权码授权流程
	// token 用于隐式授权流程
	ResponseType string `form:"response_type" binding:"required,oneof=code token"`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required,max=40"`
	// 表示客户端请求的访问权限范围
	Scope string `form:"scope" binding:"required"`
	// 用于在授权请求和授权响应之间传递状态，以防止 CSRF 攻击
	State string `form:"state"`
	// 指定授权服务器将用户重定向到的客户端应用程序的回调 URL
	RedirectURI string `form:"redirect_uri" binding:"required,uri"`
	//
	client *db.Client
}

func (q *getReq) InitTP(t *html.Authorize, rq string) {
	t.Action = fmt.Sprintf("/oauth2%s?%s", Path, rq)
	t.ClientName = *q.client.Name
	if q.client.Image != nil {
		t.ClientImage = *q.client.Image
	}
	util.CopyStruct(t, q)
	//
	if q.client.Scope == nil {
		return
	}
	scope := strings.Fields(q.Scope)
	for _, s := range strings.Fields(*q.client.Scope) {
		// 在管理接口那里已经确保是 k:v 的格式
		p := strings.Split(s, ":")
		for _, ss := range scope {
			if p[0] == ss {
				t.Scope = append(t.Scope, &html.AuthorizeScope{
					Key:   p[0],
					Name:  p[1],
					Check: true,
				})
				break
			}
		}
	}
}

// get 处理第三方授权调用
func get(ctx *gin.Context) {
	// 参数
	var req getReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorQuery, err.Error())
		return
	}
	// 应用
	req.client, err = db.GetClient(req.ClientID)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorDB, err.Error())
		return
	}
	if req.client == nil || *req.client.Enable != db.True {
		html.ExecError(ctx.Writer, html.TitleAuthorize, "第三方应用不存在", "")
		return
	}
	// 处理
	switch req.ResponseType {
	case ResponseTypeCode:
		getCode(ctx, &req)
	case ResponseTypeToken:
		getToken(ctx, &req)
	}
}

// getCode response_type=code
func getCode(ctx *gin.Context, req *getReq) {
	var tp html.Authorize
	req.InitTP(&tp, ctx.Request.URL.RawQuery)
	tp.Exec(ctx.Writer)
}

// getToken response_type=token
func getToken(ctx *gin.Context, req *getReq) {
	// var tp html.Authorize
	// req.InitTP(&tp)
	// tp.Exec(ctx.Writer)
}
