package db

import (
	"context"
	"encoding/json"
	"oauth2/cfg"
	"time"

	"github.com/qq51529210/uuid"
)

// redis key 的前缀
const (
	AccessTokenPrefix = "access_token:"
)

// 请求的类型
const (
	GrantTypeAuthorizationCode = "authorization_code"
	GrantTypeImplicit          = "implicit"
	GrantTypePassword          = "password"
	GrantTypeClientCredentials = "client_credentials"
	GrantTypeRefreshToken      = "refresh_token"
)

// AccessToken 表示访问令牌，使用 redis 来保存
type AccessToken struct {
	ID       string `json:"access_token" query:"access_token"`
	Type     string `json:"token_type" query:"token_type"`
	Expires  int64  `json:"expires_in" query:"expires_in"`
	Refresh  string `json:"refresh_token" query:"refresh_token"`
	Scope    string `json:"scope" query:"scope"`
	Grant    string `json:"grant_type" query:"grant_type"`
	UserID   string `json:"user_id" query:"user_id"`
	ClientID string `json:"client_id" query:"client_id"`
}

// PutAccessTokenWithContext 创建访问令牌
func PutAccessTokenWithContext(ctx context.Context, token *AccessToken) error {
	pip := rds.Pipeline()
	//
	token.ID = uuid.LowerV1WithoutHyphen()
	token.Expires = cfg.Cfg.OAuth2.AccessTokenExpires
	token.Refresh = uuid.LowerV1WithoutHyphen()
	data, _ := json.Marshal(token)
	err := pip.Set(ctx, AccessTokenPrefix+token.ID, data, time.Duration(cfg.Cfg.OAuth2.AccessTokenExpires)*time.Second).Err()
	if err != nil {
		return nil
	}
	//
	refreshToken := new(RefreshToken)
	refreshToken.AccessToken = *token
	refreshToken.AccessToken.ID = token.Refresh
	refreshToken.AccessToken.Refresh = uuid.LowerV1WithoutHyphen()
	refreshToken.OldAccessToken = token.ID
	data, _ = json.Marshal(refreshToken)
	err = pip.Set(ctx, RefreshTokenPrefix+refreshToken.AccessToken.ID, data, time.Duration(cfg.Cfg.OAuth2.RefreshTokenExpires)*time.Second).Err()
	if err != nil {
		return nil
	}
	//
	_, err = pip.Exec(ctx)
	return err
}

// PutAccessToken 创建访问令牌
func PutAccessToken(token *AccessToken) error {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return PutAccessTokenWithContext(ctx, token)
}

// GetAccessTokenWithContext 查询访问令牌
func GetAccessTokenWithContext(ctx context.Context, token string) (*AccessToken, error) {
	return GetWithContext[AccessToken](ctx, AccessTokenPrefix+token)
}

// GetAccessToken 查询访问令牌
func GetAccessToken(token string) (*AccessToken, error) {
	return Get[AccessToken](AccessTokenPrefix + token)
}
