package apps

import (
	"net/http"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

// @Summary  第三方应用管理
// @Tags     获取
// @Description 获取数据
// @Param    id path string true "App.ID"
// @Security ApiKeyAuth
// @Success  200 {object} db.App
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /apps/{id} [get]
func get(ctx *gin.Context) {
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session)
	// 数据库
	model, err := db.GetApp(ctx.Params[0].Value)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 没有数据，或者不是自己的
	if model == nil || model.DeveloperID != sess.User.ID {
		internal.Data404(ctx)
		return
	}
	//
	ctx.JSON(http.StatusOK, model)
}
