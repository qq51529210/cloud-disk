package authorize

import (
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
)

type getReq struct {
	// 指定用于授权流程的响应类型，常见的值包括
	// code 用于授权码授权流程
	// token 用于隐式授权流程
	ResponseType string `form:"response_type" binding:"required,oneof=code token"`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required,max=40"`
	// 表示客户端请求的访问权限范围
	Scope string `form:"scope" binding:"required,contains=image name"`
	// 用于在授权请求和授权响应之间传递状态，以防止 CSRF 攻击
	State string `form:"state"`
	// 指定授权服务器将用户重定向到的客户端应用程序的回调 URL
	RedirectURI string `form:"redirect_uri" binding:"required,uri"`
}

// get 处理第三方授权调用
func get(ctx *gin.Context) {
	// 参数
	var req getReq
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
	Action  string
}

// getResponseTypeCode 处理 response_type=code
func getResponseTypeCode(ctx *gin.Context, req *getReq) {
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
	// 返回
	var res getAuthorize
	res.AppName = *app.Name
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
