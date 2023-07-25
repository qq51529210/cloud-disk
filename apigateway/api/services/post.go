package services

import (
	"net/http"

	"gbs/api/internal"
	"gbs/api/internal/middleware"
	"gbs/db"
	"gbs/util"
	"gbs/zlm"

	"github.com/gin-gonic/gin"
)

type postReq struct {
	// 所属的流媒体服务数据库 ID
	MediaServerID *int64 `json:"mediaServerID" binding:"required,min=1"`
	// 名称，可读性
	Name *string `json:"name" binding:"required,max=32"`
	// 原始拉流地址
	SrcURL *string `json:"srcURL" binding:"required,max=255"`
	// ffmpeg 拉流的命令
	FFMPEGCmd *string `json:"ffmpegCmd" binding:"omitempty,max=32"`
	// 拉流超时，单位毫秒，默认是 10000
	Timeout *int64 `json:"timeout" binding:"omitempty,min=1000"`
	// 是否启动
	Enable *int8 `json:"enable" binding:"required,oneof=0 1"`
}

// @BaseURI
// @Summary  添加
// @Tags     拉流
// @Param    data body postReq true "数据"
// @Security ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success  201 {object} internal.IDResult[int64]
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /pull_streams [post]
func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Handle400(ctx, err.Error())
		return
	}
	ctx.Set(middleware.ReqDataCtxKey, &req)
	// 数据库
	var model db.PullStream
	util.CopyStruct(&model, &req, false)
	rows, err := db.AddPullStream(&model)
	if err != nil {
		internal.HandleDB500(ctx, err)
		return
	}
	// 内存
	if rows > 0 {
		zlm.RestartPullStream(model.ID)
	}
	// 返回
	ctx.JSON(http.StatusCreated, &internal.IDResult[string]{
		ID: model.ID,
	})
}
