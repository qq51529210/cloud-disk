package authorize

import (
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
)

// code 处理授权码模式
func code(ctx *gin.Context, req *getReq) {
	// 应用
	Client, err := db.GetClient(req.ClientID)
	if err != nil {
		errorTP.Execute(ctx.Writer, "数据库错误")
		return
	}
	if Client == nil || *Client.Enable != db.True {
		errorTP.Execute(ctx.Writer, "第三方应用不存在")
		return
	}
	// 返回
	var res getRes
	res.ClientName = *Client.Name
	res.Scope = make(map[string]string)
	for _, scope := range strings.Fields(req.Scope) {
		name, ok := authorizeName[scope]
		if ok {
			res.Scope[scope] = name
		}
	}
	res.Action = ctx.Request.URL.String()
	// 页面
	authorizeTP.Execute(ctx.Writer, &res)
}
