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
	// 账号
	Account string `json:"account" binding:"required,max=40"`
	// 密码
	Password *string `json:"password" binding:"required,max=40"`
	// 是否启用，0/1
	Enable *int8 `json:"enable" binding:"omitempty,oneof=0 1"`
}

// @Summary  添加
// @Tags     管理员
// @Param    data body postReq true "添加的数据"
// @Security ApiKeyAuth
// @Accept   json
// @Success  201
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /admins [post]
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
	var model db.Admin
	util.CopyStruct(&model, &req)
	model.ID = uuid.LowerV1WithoutHyphen()
	_, err = db.AdminDA.Add(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.Status(http.StatusCreated)
}
