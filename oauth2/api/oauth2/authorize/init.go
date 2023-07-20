package authorize

import (
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Path 是路径
const Path = "/authorize"

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group(Path, middleware.CheckUserSession)
	//
	g.GET("", get)
	g.POST("", post)
}

var (
	authorizeName = make(map[string]string)
)

func init() {
	authorizeName["image"] = "图像"
	authorizeName["name"] = "名称"
}
