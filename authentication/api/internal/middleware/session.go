package middleware

import (
	"authentication/api/internal"
	"authentication/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// CookieName 表示 cookie 的名称
	CookieName = "sid"
	// AuthenticationURL 表示用户认证的 url
	authenticationURL = "/oauth2/authentication"
	// SessionContextKey 表示 session 上下文数据的 key
	SessionContextKey = "sck"
)

// CheckSession 使用 cookie 检查用户登录
func CheckSession(ctx *gin.Context) {
	// 提取 cookie
	sid, err := ctx.Cookie(CookieName)
	if err == http.ErrNoCookie {
		ctx.Redirect(http.StatusFound, authenticationURL)
		return
	}
	// 查询 session
	sess, err := db.GetSession(sid)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 没有
	if sess == nil {
		ctx.Redirect(http.StatusFound, authenticationURL)
		return
	}
	// 设置上下文
	ctx.Set(SessionContextKey, sess)
	//
	ctx.Next()
}
