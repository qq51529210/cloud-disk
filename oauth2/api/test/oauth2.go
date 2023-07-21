package test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"oauth2/api/internal"
	"oauth2/cfg"
	"oauth2/db"

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
	// 获取 access_token
	token := getAccessToken(ctx, req.Code)
	if token == nil {
		return
	}
	// 成功
	ctx.JSON(http.StatusOK, token)
}

type oauth2TokenReq struct {
	GrantTpe     string `query:"grant_type"`
	Code         string `query:"code"`
	ClientID     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
}

func getAccessToken(ctx *gin.Context, code string) *db.AccessToken {
	// 查询参数
	var req oauth2TokenReq
	req.GrantTpe = "authorization_code"
	req.Code = code
	req.ClientID = client
	req.ClientSecret = pwd
	q := util.HTTPQuery(&req, nil)
	// 请求
	res := new(db.AccessToken)
	url := fmt.Sprintf("http://%s/oauth2/token", cfg.Cfg.Addr)
	err := util.HTTP[int](http.MethodPost, url, q, nil, res, func(res *http.Response) error {
		return util.HTTPStatusErrorHandle(res, http.StatusOK)
	}, apiCallTimeout)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			internal.Error504(ctx, err)
		} else {
			internal.Error502(ctx, err)
		}
		return nil
	}
	return res
}
