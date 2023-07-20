package authorize

import (
	"oauth2/api/internal/html"

	"github.com/gin-gonic/gin"
)

// code 处理授权码模式
func code(ctx *gin.Context, req *getReq) {
	var tp html.Authorize
	req.InitTP(&tp)
	tp.Exec(ctx.Writer)
}
