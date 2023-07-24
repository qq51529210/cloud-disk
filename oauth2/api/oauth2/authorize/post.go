package authorize

import (
	"net/http"
	"oauth2/api/internal/html"
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/log"
)

const (
	queryNameState = "state"
	queryNameCode  = "code"
)

type postReq struct {
	baseQuery
	FormID string `form:"form_id" binding:"required"`
	form   *db.AuthorizationForm
}

// 合并提交表单的值
func parsePostScope(ctx *gin.Context, client *db.Client) string {
	var scope strings.Builder
	for k := range ctx.Request.PostForm {
		if strings.Contains(*client.Scope, k) {
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

// post 处理用户确认授权
func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorQuery, err.Error())
		return
	}
	// 表单
	req.form, err = db.GetAuthorizationForm(req.FormID)
	if err != nil {
		html.ExecError(ctx.Writer, html.TitleAuthorize, html.ErrorDB, err.Error())
		return
	}
	if req.form == nil {
		ctx.Redirect(http.StatusFound, ctx.Request.URL.String())
		return
	}
	// 模式
	switch req.ResponseType {
	case ResponseTypeCode:
		postCode(ctx, &req)
	case ResponseTypeToken:
		postToken(ctx, &req)
	}
	// 删除表单
	err = db.DelAuthorizationForm(req.form.ID)
	if err != nil {
		log.Error(err)
	}
}
