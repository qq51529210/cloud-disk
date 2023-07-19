package test

import (
	"net/http"
	"oauth2/api/internal"

	"github.com/gin-gonic/gin"
)

func token(ctx *gin.Context) {
	// 参数
	var req oauth2Req
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 验证
	if req.State != state {
		internal.Submit400(ctx, "[state]不正确")
		return
	}
	// 获取 access_token
	token := getAccessToken(ctx, &req)
	if token == nil {
		return
	}
	// 成功
	ctx.JSON(http.StatusOK, "登录成功")
}
