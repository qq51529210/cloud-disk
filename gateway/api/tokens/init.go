package tokens

import (
	"gateway/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(router gin.IRouter) {
	router = router.Group("/tokens")
	//
	router.POST("", post)
	router.DELETE("", middleware.Authorization, delete)
}
