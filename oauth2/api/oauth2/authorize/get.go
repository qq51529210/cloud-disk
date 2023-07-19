package authorize

import (
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	// 参数
	var req Model
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		errorTP.Execute(ctx.Writer, "参数错误")
		return
	}
	// 类型
	switch req.ResponseType {
	case "code":
		getResponseTypeCode(ctx, &req)
	// case "token":
	default:
		errorTP.Execute(ctx.Writer, "参数错误")
	}
}

type getAuthorize struct {
	AppName string
	Scope   map[string]string
}

func getResponseTypeCode(ctx *gin.Context, req *Model) {
	// 应用
	app, err := db.GetApp(req.ClientID)
	if err != nil {
		errorTP.Execute(ctx.Writer, "数据库错误")
		return
	}
	if app == nil || *app.Enable != db.True {
		errorTP.Execute(ctx.Writer, "第三方应用不存在")
		return
	}
	// 返回授权页面
	var res getAuthorize
	res.AppName = *app.Name
	res.Scope = make(map[string]string)
	for _, scope := range strings.Fields(req.Scope) {
		name, ok := authorizeName[scope]
		if ok {
			res.Scope[scope] = name
		}
	}
	authorizeTP.Execute(ctx.Writer, &res)
}
