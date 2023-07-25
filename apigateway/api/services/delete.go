package services

import (
	"net/http"

	"gbs/api/internal"
	"gbs/api/internal/middleware"
	"gbs/db"
	"gbs/zlm"

	"github.com/gin-gonic/gin"
)

// @Summary  删除
// @Tags     拉流
// @Param    id path string true "数据库 ID"
// @Security ApiKeyAuth
// @Produce  json
// @Success  204 {object} internal.RowResult
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /pull_streams/{id} [delete]
func delete(ctx *gin.Context) {
	// 参数
	var id internal.IDPath[string]
	err := ctx.ShouldBindUri(&id)
	if err != nil {
		internal.Handle400(ctx, err.Error())
		return
	}
	// 数据库
	rows, err := db.DeletePullStream(id.ID)
	if err != nil {
		internal.HandleDB500(ctx, err)
		return
	}
	// 内存
	if rows > 0 {
		zlm.StopPullStream(id.ID)
	}
	// 返回
	ctx.JSON(http.StatusNoContent, &internal.RowResult{
		Row: rows,
	})
}

// @Summary  批量删除
// @Tags     拉流
// @Param    data body internal.BatchDelete[string] true "条件"
// @Security ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success  204 {object} internal.RowResult
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /pull_streams [delete]
func batchDelete(ctx *gin.Context) {
	// 参数
	var req internal.BatchDelete[string]
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Handle400(ctx, err.Error())
		return
	}
	ctx.Set(middleware.ReqDataCtxKey, &req)
	// 数据库
	rows, err := db.BatchDeletePullStream(req.ID)
	if err != nil {
		internal.HandleDB500(ctx, err)
		return
	}
	// 内存
	if rows > 0 {
		for _, id := range req.ID {
			zlm.StopPullStream(id)
		}
	}
	// 返回
	ctx.JSON(http.StatusNoContent, &internal.RowResult{
		Row: rows,
	})
}
