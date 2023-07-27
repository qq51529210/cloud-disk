package services

import (
	"gateway/api/internal"
	"gateway/api/internal/middleware"
	"gateway/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
	"github.com/qq51529210/uuid"
)

type postReq struct {
	// 名称，好记
	Name *string `json:"name" binding:"required,max=40"`
	// 代理路径，/order 这样的
	Path *string `json:"path" binding:"required,max=40"`
	// 是否启用，0/1
	Enable *int8 `json:"enable" binding:"required,oneof=0 1"`
}

// @Summary  添加
// @Tags     服务
// @Param    data body postReq true "添加的数据"
// @Security ApiKeyAuth
// @Accept   json
// @Success  201
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /services [post]
func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	ctx.Set(middleware.LogDataContextKey, &req)
	// 数据库
	var model db.Service
	util.CopyStruct(&model, &req)
	model.ID = uuid.LowerV1WithoutHyphen()
	_, err = db.ServiceDA.Add(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.Status(http.StatusCreated)
}
