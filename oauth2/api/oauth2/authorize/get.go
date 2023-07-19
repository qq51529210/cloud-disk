package authorize

import (
	"oauth2/db"
	"strings"

	"github.com/gin-gonic/gin"
)

// 模式
const (
	ResponseTypeCode              = "code"
	ResponseTypeToken             = "token"
	ResponseTypePassword          = "password"
	ResponseTypeClientCredentials = "client_credentials"
)

const (
	errQuery = "参数错误"
)

var (
	responseTypeHandle = make(map[string]func(*gin.Context, *getReq))
)

func init() {
	responseTypeHandle[ResponseTypeCode] = code
	responseTypeHandle[ResponseTypeToken] = token
	responseTypeHandle[ResponseTypePassword] = password
	responseTypeHandle[ResponseTypeClientCredentials] = clientCredentials
}

type getReq struct {
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
	//
	client *db.Client
}

type getRes struct {
	// 应用名称
	ClientName string
	// 应用图片
	ClientImage string
	// 访问权限
	Scope map[string]string
	// 表单地址
	Action string
}

func (s *getRes) Init(q *getReq) {
	image := ""
	if q.client.Image != nil {
		image = *q.client.Image
	}
	s.ClientImage = image
	s.ClientName = *q.client.Name
	s.Scope = make(map[string]string)
	for _, scope := range strings.Fields(q.Scope) {
		name, ok := authorizeName[scope]
		if ok {
			s.Scope[scope] = name
		}
	}
}

// get 处理第三方授权调用
func get(ctx *gin.Context) {
	// 参数
	var req getReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		errorTP.Execute(ctx.Writer, errQuery)
		return
	}
	// 应用
	req.client, err = db.GetClient(req.ClientID)
	if err != nil {
		errorTP.Execute(ctx.Writer, "数据库错误")
		return
	}
	if req.client == nil || *req.client.Enable != db.True {
		errorTP.Execute(ctx.Writer, "第三方应用不存在")
		return
	}
	// 处理
	hd, ok := responseTypeHandle[req.ResponseType]
	if !ok {
		errorTP.Execute(ctx.Writer, errQuery)
		return
	}
	hd(ctx, &req)
}
