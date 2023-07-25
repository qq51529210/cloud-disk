package services

import (
	"net/http"

	"gbs/api/internal"
	"gbs/api/internal/middleware"
	"gbs/db"
	"gbs/zlm"

	"github.com/gin-gonic/gin"
)

// @Summary  列表
// @Tags     拉流
// @Param    query query db.PullStreamQuery false "条件"
// @Security ApiKeyAuth
// @Produce  json
// @Success  200 {object} db.ListData[db.PullStream]
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /pull_streams [get]
func list(ctx *gin.Context) {
	// 参数
	var req db.PullStreamQuery
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Handle400(ctx, err.Error())
		return
	}
	ctx.Set(middleware.ReqDataCtxKey, &req)
	// 数据库
	var res db.ListData[*db.PullStream]
	err = db.List(&req, &req.Page, &res)
	if err != nil {
		internal.HandleDB500(ctx, err)
		return
	}
	// 流状态
	zlm.BatchSetupPullStream(res.Data)
	// 返回
	ctx.JSON(http.StatusOK, &res)
}
