package authorize

import (
	"oauth2/api/internal/html"

	"github.com/gin-gonic/gin"
)

// token 处理隐式模式
func token(ctx *gin.Context, req *getReq) {
	var tp html.Authorize
	req.InitTP(&tp)
	tp.Exec(ctx.Writer)
}
