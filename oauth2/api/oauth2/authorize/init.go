package authorize

import (
	"oauth2/api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Path 是路径
const Path = "/authorize"

// Init 初始化路由
func Init(g gin.IRouter) {
	g = g.Group(Path, middleware.CheckSession)
	//
	g.GET("", get)
	g.POST("", post)
}

// Model 表示第三方跳转到 GET /authorize 的查询参数
type Model struct {
	// 指定用于授权流程的响应类型，常见的值包括
	// code 用于授权码授权流程
	// token 用于隐式授权流程
	ResponseType string `form:"response_type" binding:"required,oneof=code token"`
	// 表示客户端应用程序的唯一标识符，由授权服务器分配给客户端
	ClientID string `form:"client_id" binding:"required,max=40"`
	// 表示客户端请求的访问权限范围
	Scope string `form:"scope" binding:"required,contains=image name"`
	// 用于在授权请求和授权响应之间传递状态，以防止 CSRF 攻击
	State string `form:"state"`
	// 指定授权服务器将用户重定向到的客户端应用程序的回调 URL
	RedirectURI string `form:"redirect_uri" binding:"required,uri"`
}

var (
	authorizeName = make(map[string]string)
)

func init() {
	authorizeName["image"] = "图像"
	authorizeName["name"] = "名称"
}
