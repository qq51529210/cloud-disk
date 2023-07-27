package tokens

import (
	"gateway/api/internal/middleware"
	"gateway/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 删除
// @Tags    Token
// @Produce json
// @Success 204
// @Failure 401
// @Failure 500 {object} internal.Error
// @Router  /tokens [delete]
func delete(ctx *gin.Context) {
	// 数据库
	db.DelSession(ctx.Value(middleware.TokenContextKey).(string))
	// 返回
	ctx.Status(http.StatusNoContent)
}
