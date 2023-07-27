package servers

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
	// 服务数据库 id
	ServiceID string `json:"serviceID" binding:"required,max=40"`
	// 基本路径，http(https)://hostname:port/
	BaseURL string `json:"baseURL" binding:"required,max=128,url"`
	// 名称，好记
	Name *string `json:"name" binding:"required,max=40"`
	// 是否启用，0/1
	Enable *int8 `json:"enable" binding:"omitempty,oneof=0 1"`
	// 开启身份认证，0/1
	Authorization *int8 `json:"authorization" binding:"omitempty,oneof=0 1"`
	// 访问控制，单位，次/每秒
	Limite *int32 `json:"limite" binding:"omitempty,min=0"`
}

// @Summary  添加
// @Tags     服务器
// @Param    data body postReq true "添加的数据"
// @Security ApiKeyAuth
// @Accept   json
// @Success  201
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /servers [post]
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
	var model db.Server
	util.CopyStruct(&model, &req)
	model.ID = uuid.LowerV1WithoutHyphen()
	_, err = db.ServerDA.Add(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.Status(http.StatusCreated)
}
