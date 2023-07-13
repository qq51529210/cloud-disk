package users

import (
	"authentication/api/internal"
	"authentication/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
	"github.com/qq51529210/uuid"
)

type postReq struct {
	// 账号
	Account *string `json:"account" binding:"required,max=32"`
	// 密码，SHA1 格式
	Password *string `json:"password" binding:"required"`
}

// @Summary  用户管理
// @Tags     添加
// @Param    data body postReq true "添加的字段"
// @Security ApiKeyAuth
// @Success  201
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /users [post]
func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 数据库
	var model db.User
	util.CopyStructAll(&model, &req)
	model.ID = uuid.LowerV1WithoutHyphen()
	_, err = db.AddUser(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	//
	ctx.Status(http.StatusCreated)
}
