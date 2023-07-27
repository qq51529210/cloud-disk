package services

import (
	"gateway/api/internal"
	"gateway/api/internal/middleware"
	"gateway/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

// @Summary  列表
// @Tags     服务
// @Param    query query db.ServiceQuery false "条件"
// @Security ApiKeyAuth
// @Produce  json
// @Success  200 {object} util.GORMList[db.Service]
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /services [get]
func list(ctx *gin.Context) {
	// 参数
	var req db.ServiceQuery
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	ctx.Set(middleware.LogDataContextKey, &req)
	// 数据库
	var res util.GORMList[*db.Service]
	err = db.ServiceDA.List(&req.GORMPage, &req, &res)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.JSON(http.StatusOK, &res)
}
