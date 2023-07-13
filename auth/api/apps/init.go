package apps

import (
	"auth/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/apps", middleware.CheckSession)
	//
	g.GET("", get)
	g.POST("", post)
	g.PATCH("/:id", patch)
	g.DELETE("/:id", delete)
}
