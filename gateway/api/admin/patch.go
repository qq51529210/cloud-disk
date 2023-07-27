package services

import (
	"gateway/api/internal"
	"gateway/api/internal/middleware"
	"gateway/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

type patchReq struct {
	// 密码
	Password *string `json:"password" binding:"omitempty,max=40"`
	// 是否启用，0/1
	Enable *int8 `json:"enable" binding:"omitempty,oneof=0 1"`
}

// @Summary  修改
// @Tags     管理员
// @Param    id path string true "id"
// @Param    data body patchReq true "修改的数据"
// @Security ApiKeyAuth
// @Accept   json
// @Success  200
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  500 {object} internal.Error
// @Router   /admins/{id} [patch]
func patch(ctx *gin.Context) {
	// 参数
	var req patchReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	if util.IsNilOrEmpty(&req) {
		internal.SubmitEmpty400(ctx)
		return
	}
	ctx.Set(middleware.LogDataContextKey, &req)
	// 数据库
	var model db.Admin
	util.CopyStruct(&model, &req)
	model.ID = ctx.Params[0].Value
	_, err = db.AdminDA.Update(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.Status(http.StatusOK)
}
