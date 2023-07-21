package authorize

import (
	"fmt"
	"oauth2/api/internal/html"
	"oauth2/cfg"
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
	baseQuery
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
		if !cfg.Cfg.OAuth2.EnableImplicitGrant {
			html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorQuery, "不支持 implicit 模式")
			return
		}
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
	var tp html.Authorize
	req.InitTP(&tp, ctx.Request.URL.RawQuery)
	tp.Exec(ctx.Writer)
}
