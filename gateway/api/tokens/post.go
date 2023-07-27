package tokens

import (
	"gateway/api/internal"
	"gateway/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

const (
	errorPhrase = "create token fail"
)

type postReq struct {
	// 账号
	Account string `json:"account" binding:"required,max=40"`
	// 密码
	Password string `json:"password" binding:"required"`
}

type postRes struct {
	// 令牌
	Token string `json:"token"`
	// 过期时间
	Expired int64 `json:"expiredAt"`
}

// @Summary 创建
// @Tags    Token
// @Param   data body postReq true "数据"
// @Accept  json
// @Produce json
// @Success 201 {object} postRes
// @Failure 400 {object} internal.Error
// @Failure 401 {object} internal.Error
// @Failure 403 {object} internal.Error
// @Failure 500 {object} internal.Error
// @Router  /tokens [post]
func post(ctx *gin.Context) {
	// 参数
	var req postReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 数据库
	var model db.Admin
	model.Account = req.Account
	ok, err := db.AdminDA.Get(&model)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 不存在/错误
	if !ok || model.Password != util.SHA1String(req.Password) {
		ctx.JSON(http.StatusUnauthorized, &internal.Error{
			Phrase: errorPhrase,
			Detail: "账号/密码错误",
		})
		return
	}
	// 禁止
	if *model.Enable != db.True {
		ctx.JSON(http.StatusForbidden, &internal.Error{
			Phrase: errorPhrase,
			Detail: "账号已禁用",
		})
		return
	}
	// 会话
	sess, err := db.NewSession()
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 返回
	ctx.JSON(http.StatusCreated, &postRes{
		Token:   sess.Token,
		Expired: sess.Expired,
	})
}
