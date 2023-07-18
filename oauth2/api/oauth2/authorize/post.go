package authorize

import (
	"auth/api/internal"
	"auth/db"

	"github.com/gin-gonic/gin"
)

func post(ctx *gin.Context) {
	// 参数
	var req reqQuery
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 查询
	app, err := db.GetApp(req.ClientID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	if app == nil {
		internal.Data404(ctx)
		return
	}
	// 跳转
}
