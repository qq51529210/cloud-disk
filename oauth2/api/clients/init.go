package clients

import (
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group("/clients", middleware.CheckDeveloperSession)
	//
	g.GET("", get)
	g.POST("", post)
	g.PATCH("/:id", patch)
	g.DELETE("/:id", delete)
}
