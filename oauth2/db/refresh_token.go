package db

import (
	"context"
)

// redis key 的前缀
const (
	RefreshTokenPrefix = "refresh_token:"
)

// RefreshToken 表示访问令牌，使用 redis 来保存
type RefreshToken struct {
	AccessToken
	OldAccessToken string
}

// GetRefreshTokenWithContext 查询刷新令牌
func GetRefreshTokenWithContext(ctx context.Context, token string) (*RefreshToken, error) {
	return GetWithContext[RefreshToken](ctx, RefreshTokenPrefix+token)
}

// GetRefreshToken 查询刷新令牌
func GetRefreshToken(token string) (*RefreshToken, error) {
	return Get[RefreshToken](RefreshTokenPrefix + token)
}

// DelRefreshTokenWithContext 查询授权码
func DelRefreshTokenWithContext(ctx context.Context, code *RefreshToken) error {
	pip := rds.Pipeline()
	//
	err := pip.Del(ctx, RefreshTokenPrefix+code.AccessToken.ID).Err()
	if err != nil {
		return nil
	}
	err = pip.Del(ctx, AccessTokenPrefix+code.OldAccessToken).Err()
	if err != nil {
		return nil
	}
	//
	_, err = pip.Exec(ctx)
	return err
}

// DelRefreshToken 查询授权码
func DelRefreshToken(code *RefreshToken) error {
	ctx, cancel := newRedisTimeout()
	defer cancel()
	return DelRefreshTokenWithContext(ctx, code)
}
