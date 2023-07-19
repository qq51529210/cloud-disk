package clients

import (
	"net/http"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

type patchReq struct {
	// 密码，SHA1 格式
	Secret *string `json:"secret" binding:"omitempty,max=40"`
	// 名称
	Name *string `json:"name" binding:"omitempty,max=40"`
	// 描述
	Description *string `json:"description" binding:"omitempty,max=255"`
	// 是否启用，0/1
	Enable *int8 `json:"enable" binding:"omitempty,oneof=0 1"`
	// 重定向 url 列表
	URLs []string `json:"urls" binding:"omitempty,dive,url"`
}

// @Summary  第三方应用
// @Tags     修改
// @Description 修改数据
// @Param    id path string true "Client.ID"
// @Param    data body patchReq true "修改的字段"
// @Security ApiKeyAuth
// @Success  201
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /clients/{id} [patch]
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
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session)
	// 数据库
	var model db.Client
	util.CopyStructAll(&model, &req)
	model.ID = ctx.Params[0].Value
	_, err = db.UpdateClient(&model, sess.User.ID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	//
	ctx.Status(http.StatusCreated)
}
