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
	AccessTokenPrefix  = "access_token:"
	RefreshTokenPrefix = "refresh_token:"
)

// 请求的类型
const (
	GenTypeToken       = "token"
	GenTypeCode        = "code"
	GenTypePassword    = "password"
	GenTypeRefresh     = "refresh"
	GenTypeCredentials = "credentials"
)

// AccessToken 表示访问令牌，使用 redis 来保存
type AccessToken struct {
	ID       string
	Type     string
	Expires  int64
	Refresh  string
	Scope    string
	GenType  string
	UserID   string
	ClientID string
}

// PutAccessTokenWithContext 创建访问令牌
func PutAccessTokenWithContext(ctx context.Context, token *AccessToken) error {
	token.ID = uuid.LowerV1WithoutHyphen()
	token.Refresh = uuid.LowerV1WithoutHyphen()
	token.Expires = cfg.Cfg.OAuth2.AccessTokenExpires
	pip := rds.Pipeline()
	data, _ := json.Marshal(token)
	err := pip.Set(ctx, AccessTokenPrefix+token.ID, data, time.Duration(cfg.Cfg.OAuth2.AccessTokenExpires)*time.Second).Err()
	if err != nil {
		return nil
	}
	//
	refreshToken := new(AccessToken)
	*refreshToken = *token
	refreshToken.ID = token.Refresh
	refreshToken.Refresh = uuid.LowerV1WithoutHyphen()
	data, _ = json.Marshal(token)
	err = pip.Set(ctx, RefreshTokenPrefix+token.ID, data, time.Duration(cfg.Cfg.OAuth2.RefreshTokenExpires)*time.Second).Err()
	if err != nil {
		return nil
	}
	//
	_, err = pip.Exec(ctx)
	//
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

// GetRefreshTokenWithContext 查询刷新令牌
func GetRefreshTokenWithContext(ctx context.Context, token string) (*AccessToken, error) {
	return GetWithContext[AccessToken](ctx, RefreshTokenPrefix+token)
}

// GetRefreshToken 查询刷新令牌
func GetRefreshToken(token string) (*AccessToken, error) {
	return Get[AccessToken](RefreshTokenPrefix + token)
}
