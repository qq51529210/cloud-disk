package users

// import (
// 	"auth/api/internal"
// 	"auth/api/internal/middleware"
// 	"auth/db"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// // @Summary  用户管理
// // @Tags     删除
// // @Description 删除数据
// // @Param    id path string true "User.ID"
// // @Security ApiKeyAuth
// // @Success  204
// // @Failure  400 {object} internal.Error
// // @Failure  401
// // @Failure  403
// // @Failure  500 {object} internal.Error
// // @Router   /users/{id} [delete]
// func delete(ctx *gin.Context) {
// 	// 会话
// 	sess := ctx.Value(middleware.SessionContextKey).(*db.Session)
// 	// 数据库
// 	_, err := db.DeleteUser(ctx.Params[0].Value, sess.User.ID)
// 	if err != nil {
// 		internal.DB500(ctx, err)
// 		return
// 	}
// 	//
// 	ctx.Status(http.StatusNoContent)
// }
