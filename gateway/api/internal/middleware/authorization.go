package middleware

import (
	"gateway/db"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// TokenContextKey 是 gin.Context 的 token 上下文 key
	TokenContextKey = "token"
)

// Authorization 拦截并检查 Authorization: Bearer token
func Authorization(ctx *gin.Context) {
	// token
	token := bearerToken(ctx)
	if token == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// session
	sess, err := db.GetSession(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if sess == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// 通过
	ctx.Set(TokenContextKey, sess)
}

// bearerToken 解析并返回 Authorization 头的 bearer token
func bearerToken(ctx *gin.Context) string {
	authorizationHeader := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	if token == authorizationHeader {
		return ""
	}
	return token
}
