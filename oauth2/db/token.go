package db

import (
	"context"
	"oauth2/cfg"
	"time"

	"github.com/qq51529210/uuid"
)

// AccessToken 表示访问令牌，使用 redis 来保存
type AccessToken struct {
	// uuid
	ID string
	// 过期时间
	Expires int64
	// 应用 id
	ClientID string
	// 用户 id
	UserID string
	// 范围
	Scope string
}

// NewAccessToken 创建新的访问令牌
func NewAccessToken(user *User) (*AccessToken, error) {
	return nil, nil
}

// GetAccessToken 获取会话
func GetAccessToken(id string) (*AccessToken, error) {
	return nil, nil
}

// NewAuthorizationCode 创建新的授权码
func NewAuthorizationCode(ctx context.Context) (string, error) {
	code := uuid.LowerV1WithoutHyphen()
	err := rds.Set(ctx, code, "", time.Duration(cfg.Cfg.OAuth2.AuthorizationCodeExpires)*time.Second).Err()
	if err != nil {
		return "", err
	}
	return code, nil
}

// NewAuthorizationCodeTimeout 创建授权码
func NewAuthorizationCodeTimeout() (string, error) {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return NewAuthorizationCode(ctx)
}

// GetAuthorizationCode 查询授权码
func GetAuthorizationCode(ctx context.Context, code string) (bool, error) {
	n, err := rds.Exists(ctx, code).Result()
	if err != nil {
		return false, err
	}
	return n == 1, nil
}

// GetAuthorizationCodeTimeout 查询授权码
func GetAuthorizationCodeTimeout(code string) (bool, error) {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return GetAuthorizationCode(ctx, code)
}
