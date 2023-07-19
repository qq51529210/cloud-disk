package apps

import (
	"net/http"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
	"github.com/qq51529210/uuid"
)

type postReq struct {
	// 密码，SHA1 格式
	Secret *string `json:"secret" binding:"required,max=40"`
	// 名称
	Name *string `json:"name" binding:"required,max=40"`
	// 描述
	Description *string `json:"description" binding:"omitempty,max=255"`
	// 是否启用，0/1
	Enable *int8 `json:"enable" binding:"omitempty,oneof=0 1"`
	// 重定向 url 列表
	URLs []string `json:"urls" binding:"required,dive,url"`
}

// @Summary  第三方应用管理
// @Tags     添加
// @Param    data body postReq true "添加的字段"
// @Security ApiKeyAuth
// @Success  201
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /apps [post]
func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session)
	// 数据库
	var model db.App
	util.CopyStructAll(&model, &req)
	model.ID = uuid.LowerV1WithoutHyphen()
	model.DeveloperID = sess.User.ID
	_, err = db.AddApp(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	//
	ctx.Status(http.StatusCreated)
}
