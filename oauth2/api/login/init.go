package login

import (
	"github.com/gin-gonic/gin"
)

// Path 是路径
const Path = "/login"

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group(Path)
	//
	g.GET("", get)
	g.POST("", post)
}
