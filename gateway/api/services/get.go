package services

import (
	"gateway/api/internal"
	"gateway/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary  详情
// @Tags     服务
// @Param    id path string true "id"
// @Security ApiKeyAuth
// @Produce  json
// @Success  200 {object} db.Service
// @Failure  401
// @Failure  404
// @Failure  500 {object} internal.Error
// @Router   /services/{id} [get]
func get(ctx *gin.Context) {
	// 数据库
	var model db.Service
	model.ID = ctx.Params[0].Value
	ok, err := db.ServiceDA.Get(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	if ok {
		ctx.JSON(http.StatusOK, &model)
		return
	}
	ctx.Status(http.StatusNotFound)
}
