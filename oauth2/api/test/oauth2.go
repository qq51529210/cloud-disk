package test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"oauth2/api/internal"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

type oauth2AuthorizeReq struct {
	// 用于在授权请求和授权响应之间传递状态，以防止 CSRF 攻击
	State string `form:"state"`
	// 授权码，用于获取 access_token
	Code string `form:"code"`
}

func oauth2(ctx *gin.Context) {
	// 参数
	var req oauth2AuthorizeReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		internal.Submit400(ctx, err.Error())
		return
	}
	// 验证
	if req.State != state {
		internal.Submit400(ctx, "[state]不正确")
		return
	}
	// 获取 access_token
	token := getAccessToken(ctx, req.Code)
	if token == nil {
		return
	}
	// 成功
	ctx.JSON(http.StatusOK, "登录成功")
}

type oauth2TokenReq struct {
	GrantTpe     string `query:"grant_type"`
	Code         string `query:"code"`
	ClientID     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
}

type oauth2TokenRes struct {
	// 应用程序在请求访问受保护资源时使用的令牌。
	// 它代表了客户端被授权的权限
	AccessToken string `json:"AccessToken"`
	// 该字段指示返回的令牌类型。
	// 比如 Bearer 令牌，意味着客户端可以简单地在后续请求的 "oauth2orization" 头中附上该令牌
	TokenType string `json:"token_type"`
	// 该字段以秒为单位指定访问令牌的过期时间。
	// 在此时间之后，访问令牌将不再有效，客户端需要获取新的访问令牌。
	Expires int64 `json:"expires_in"`
	// 该令牌可由客户端用于在当前访问令牌过期时获取新的访问令牌。
	// 通常在OAuth2的刷新令牌流程中使用，
	// 以便在不需要用户重新认证的情况下获取新的访问令牌
	RefreshToken string `json:"refresh_token"`
}

func getAccessToken(ctx *gin.Context, code string) *oauth2TokenRes {
	// 查询参数
	var req oauth2TokenReq
	req.GrantTpe = "authorization_code"
	req.Code = code
	req.ClientID = client
	req.ClientSecret = pwd
	q := util.HTTPQuery(&req)
	// 请求
	var res oauth2TokenRes
	url := fmt.Sprintf("%s/oauth2/token", oauth2Host)
	err := util.HTTP[int](http.MethodPost, url, q, nil, &res, http.StatusOK, apiCallTimeout)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			internal.Error504(ctx, err)
		} else {
			internal.Error502(ctx, err)
		}
		return nil
	}
	return &res
}
