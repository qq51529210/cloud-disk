package servers

import (
	"net/http"

	"gateway/api/internal"
	"gateway/api/internal/middleware"
	"gateway/db"

	"github.com/gin-gonic/gin"
)

// @Summary  删除
// @Tags     服务器
// @Param    id path string true "数据库 ID"
// @Security ApiKeyAuth
// @Success  204
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /servers/{id} [delete]
func delete(ctx *gin.Context) {
	// 数据库
	_, err := db.ServerDA.Delete(ctx.Params[0].Value)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.Status(http.StatusNoContent)
}

// @Summary  批量删除
// @Tags     服务器
// @Param    id body []string true "id数组"
// @Security ApiKeyAuth
// @Accept   json
// @Success  204
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /servers [delete]
func batchDelete(ctx *gin.Context) {
	// 参数
	var req []string
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	ctx.Set(middleware.LogDataContextKey, &req)
	// 数据库
	_, err = db.ServerDA.BatchDelete(req)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.Status(http.StatusNoContent)
}
