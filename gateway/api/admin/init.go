package services

import (
	"gateway/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(router gin.IRouter) {
	router = router.Group("/admins", middleware.Authorization)
	//
	router.GET("", list)
	router.POST("", post)
	router.DELETE("", batchDelete)
	router.GET("/:id", get)
	router.PATCH("/:id", patch)
	router.DELETE("/:id", delete)
}
