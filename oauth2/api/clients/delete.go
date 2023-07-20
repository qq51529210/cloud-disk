package clients

import (
	"net/http"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

// @Summary  第三方应用
// @Tags     删除
// @Description 删除数据
// @Param    id path string true "Client.ID"
// @Security ApiKeyAuth
// @Success  204
// @Failure  400 {object} internal.Error
// @Failure  401
// @Failure  403
// @Failure  500 {object} internal.Error
// @Router   /clients/{id} [delete]
func delete(ctx *gin.Context) {
	// 会话
	sess := ctx.Value(middleware.SessionContextKey).(*db.Session[*db.Developer])
	// 数据库
	_, err := db.DeleteClient(ctx.Params[0].Value, sess.Data.ID)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	//
	ctx.Status(http.StatusNoContent)
}
