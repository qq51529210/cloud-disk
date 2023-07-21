package db

import (
	"context"
	"oauth2/cfg"

	"github.com/qq51529210/uuid"
)

// redis key 的前缀
const (
	AccessTokenPrefix = "access_token:"
)

// AccessToken 表示访问令牌，使用 redis 来保存
type AccessToken struct {
	ID       string `json:"access_token" query:"access_token"`
	Type     string `json:"token_type" query:"token_type"`
	Expires  int64  `json:"expires_in" query:"expires_in"`
	Refresh  string `json:"refresh_token" query:"refresh_token"`
	Scope    string `json:"scope" query:"scope"`
	UserID   string `json:"user_id" query:"user_id"`
	ClientID string `json:"-"`
}

// PutAccessTokenWithContext 创建授权码
func PutAccessTokenWithContext(ctx context.Context, token *AccessToken) error {
	token.ID = uuid.LowerV1WithoutHyphen()
	token.Refresh = uuid.LowerV1WithoutHyphen()
	token.Type = "Bearer"
	token.Expires = cfg.Cfg.OAuth2.AccessTokenExpires
	return PutWithContext(ctx, AccessTokenPrefix+token.ID, token, cfg.Cfg.OAuth2.AccessTokenExpires)
}

// PutAccessToken 创建授权码
func PutAccessToken(token *AccessToken) error {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return PutAccessTokenWithContext(ctx, token)
}

// GetAccessTokenWithContext 查询授权码
func GetAccessTokenWithContext(ctx context.Context, token string) (*AccessToken, error) {
	return GetWithContext[AccessToken](ctx, AccessTokenPrefix+token)
}

// GetAccessToken 查询授权码
func GetAccessToken(token string) (*AccessToken, error) {
	return Get[AccessToken](AccessTokenPrefix + token)
}
