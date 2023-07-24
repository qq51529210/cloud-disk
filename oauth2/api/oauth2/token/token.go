package token

import (
	"fmt"
	"net/http"
	"net/url"
	"oauth2/api/internal"
	"oauth2/api/internal/middleware"
	"oauth2/cfg"
	"oauth2/db"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

// token 处理获取访问令牌
func token(ctx *gin.Context) {
	grantType := ctx.Query("grant_type")
	switch grantType {
	case db.GrantTypeAuthorizationCode:
		tokenAuthorizationCode(ctx)
	case db.GrantTypePassword:
		if cfg.Cfg.OAuth2.EnablePasswordGrant {
			tokenPassword(ctx)
		}
	case db.GrantTypeClientCredentials:
		if cfg.Cfg.OAuth2.EnableClientCredentialsGrant {
			tokenClientCredentials(ctx)
		}
	case db.GrantTypeRefreshToken:
		tokenRefreshToken(ctx)
	}
	internal.Submit400(ctx, fmt.Sprintf("[grant_type]不支持[%s]", grantType))
}

func onOK(ctx *gin.Context, token *db.Token) {
	// 重定向
	redirectURI := ctx.Query(middleware.QueryRedirectURI)
	if redirectURI != "" {
		// 重定向地址
		_u, err := url.Parse(redirectURI)
		if err != nil {
			internal.Submit400(ctx, err.Error())
			return
		}
		_u.RawQuery = util.HTTPQuery(token, _u.Query()).Encode()
		// 跳转
		ctx.Redirect(http.StatusSeeOther, _u.String())
		return
	}
	// 没有重定向，返回 JSON
	ctx.JSON(http.StatusOK, token)
}
