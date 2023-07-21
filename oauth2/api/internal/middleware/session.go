package middleware

import (
	"net/http"
	"net/url"
	"oauth2/api/internal"
	"oauth2/db"

	"github.com/gin-gonic/gin"
)

const (
	// CookieName 表示 cookie 的名称
	CookieName = "sid"
	// SessionContextKey 表示 session 上下文数据的 key
	SessionContextKey = "sck"
	// 登录 url
	loginURL = "/login"
	// QueryRedirectURI 重定向，查询参数名称
	QueryRedirectURI = "redirect_uri"
)

// CheckUserSession 使用 cookie 检查用户登录
func CheckUserSession(ctx *gin.Context) {
	// 提取 cookie
	sid, err := ctx.Cookie(CookieName)
	if err == http.ErrNoCookie {
		redirectLogin(ctx)
		return
	}
	// 查询 session
	sess, err := db.GetUserSession(sid)
	if err != nil {
		internal.DB500(ctx, err)
		return
	}
	// 没有
	if sess == nil {
		redirectLogin(ctx)
		return
	}
	// 设置上下文
	ctx.Set(SessionContextKey, sess)
	//
	ctx.Next()
}

// redirectLogin 重定向到 /login
func redirectLogin(ctx *gin.Context) {
	query := make(url.Values)
	query.Set(QueryRedirectURI, ctx.Request.URL.String())
	redirectURL := loginURL + "?" + query.Encode()
	ctx.Redirect(http.StatusFound, redirectURL)
	ctx.Abort()
}

// CheckDeveloperSession 使用 cookie 检查用户登录
func CheckDeveloperSession(ctx *gin.Context) {

}
