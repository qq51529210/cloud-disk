package authorize

import (
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

const (
	// Path 是路径
	Path = "/authorize"
)

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group(Path, middleware.CheckUserSession)
	//
	g.GET("", get)
	g.POST("", post)
}

type baseQuery struct {
	ResponseType string `form:"response_type" binding:"required,oneof=code token"`
	ClientID     string `form:"client_id" binding:"required,max=40"`
	Scope        string `form:"scope" binding:"required"`
	State        string `form:"state"`
	RedirectURI  string `form:"redirect_uri" binding:"required,uri"`
}
