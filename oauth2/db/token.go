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
	GrantTypeAuthorizationCode = "authorization_code"
	GrantTypeImplicit          = "implicit"
	GrantTypePassword          = "password"
	GrantTypeClientCredentials = "client_credentials"
	GrantTypeRefreshToken      = "refresh_token"
)

// Token 表示访问令牌，使用 redis 来保存
type Token struct {
	AccessToken  string `json:"access_token" query:"access_token"`
	TokenType    string `json:"token_type" query:"token_type"`
	ExpiresIN    int64  `json:"expires_in" query:"expires_in"`
	RefreshToken string `json:"refresh_token" query:"refresh_token"`
	Scope        string `json:"scope" query:"scope"`
	GrantType    string `json:"grant_type" query:"grant_type"`
	UserID       string `json:"user_id" query:"user_id"`
	ClientID     string `json:"client_id" query:"client_id"`
}

// PutAccessTokenWithContext 创建访问令牌
func PutAccessTokenWithContext(ctx context.Context, token *Token) error {
	token.AccessToken = uuid.LowerV1WithoutHyphen()
	token.ExpiresIN = cfg.Cfg.OAuth2.AccessTokenExpires
	token.RefreshToken = uuid.LowerV1WithoutHyphen()
	pip := rds.Pipeline()
	data, _ := json.Marshal(token)
	err := pip.Set(ctx, AccessTokenPrefix+token.AccessToken, data, time.Duration(cfg.Cfg.OAuth2.AccessTokenExpires)*time.Second).Err()
	if err != nil {
		return nil
	}
	//
	refreshToken := new(Token)
	*refreshToken = *token
	refreshToken.AccessToken = token.RefreshToken
	refreshToken.RefreshToken = uuid.LowerV1WithoutHyphen()
	data, _ = json.Marshal(token)
	err = pip.Set(ctx, RefreshTokenPrefix+token.AccessToken, data, time.Duration(cfg.Cfg.OAuth2.RefreshTokenExpires)*time.Second).Err()
	if err != nil {
		return nil
	}
	//
	_, err = pip.Exec(ctx)
	//
	return err
}

// PutAccessToken 创建访问令牌
func PutAccessToken(token *Token) error {
	// 超时
	ctx, cancel := newRedisTimeout()
	defer cancel()
	//
	return PutAccessTokenWithContext(ctx, token)
}

// GetAccessTokenWithContext 查询访问令牌
func GetAccessTokenWithContext(ctx context.Context, token string) (*Token, error) {
	return GetWithContext[Token](ctx, AccessTokenPrefix+token)
}

// GetAccessToken 查询访问令牌
func GetAccessToken(token string) (*Token, error) {
	return Get[Token](AccessTokenPrefix + token)
}

// GetRefreshTokenWithContext 查询刷新令牌
func GetRefreshTokenWithContext(ctx context.Context, token string) (*Token, error) {
	return GetWithContext[Token](ctx, RefreshTokenPrefix+token)
}

// GetRefreshToken 查询刷新令牌
func GetRefreshToken(token string) (*Token, error) {
	return Get[Token](RefreshTokenPrefix + token)
}

// DelRefreshTokenWithContext 查询授权码
func DelRefreshTokenWithContext(ctx context.Context, code string) error {
	return rds.Del(ctx, RefreshTokenPrefix+code).Err()
}

// DelRefreshToken 查询授权码
func DelRefreshToken(code string) error {
	return Del(RefreshTokenPrefix + code)
}
