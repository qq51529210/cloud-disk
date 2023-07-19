package authorize

import (
	"github.com/gin-gonic/gin"
)

// code 处理授权码模式
func code(ctx *gin.Context, req *getReq) {
	var res getRes
	res.Init(req)
	res.Action = ctx.Request.URL.String()
	authorizeTP.Execute(ctx.Writer, &res)
}
