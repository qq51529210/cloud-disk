package authorize

import (
	"github.com/gin-gonic/gin"
)

// token 处理隐式模式
func token(ctx *gin.Context, req *getReq) {
	var res getRes
	res.Init(req)
	res.Action = ctx.Request.URL.String()
	authorizeTP.Execute(ctx.Writer, &res)
}
