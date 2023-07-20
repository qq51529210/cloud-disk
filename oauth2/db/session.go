package db

import (
	"context"
	"encoding/json"
	"oauth2/cfg"
	"time"

	"github.com/qq51529210/uuid"
	"github.com/redis/go-redis/v9"
)

// redis key 的前缀
const (
	UserSessionPrefix      = "user_session_"
	DeveloperSessionPrefix = "developer_session_"
)

// Session 表示会话
type Session[T any] struct {
	// 会话 ID
	ID string
	// 用户
	Data T
	// 创建时间
	Time int64
	// 过期时间
	Expires int64
}

// NewSession 创建会话
func NewSession[T any](ctx context.Context, prefixKey string, data T) (*Session[T], error) {
	s := &Session[T]{
		ID:      prefixKey + uuid.LowerV1WithoutHyphen(),
		Data:    data,
		Time:    time.Now().Unix(),
		Expires: cfg.Cfg.Session.Expires,
	}
	// 格式化
	d, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	// 保存
	err = rds.Set(ctx, s.ID, d, time.Duration(cfg.Cfg.Session.Expires)*time.Second).Err()
	if err != nil {
		return nil, err
	}
	//
	return s, nil
}

// GetSession 获取会话
func GetSession[T any](ctx context.Context, id string) (*Session[T], error) {
	// 获取
	d, err := rds.Get(ctx, id).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	// 解析
	s := new(Session[T])
	err = json.Unmarshal(d, s)
	if err != nil {
		return nil, err
	}
	//
	return s, nil
}

// NewSessionTimeout 创建会话
func NewSessionTimeout[T any](prefixKey string, data T) (*Session[T], error) {
	// 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Cfg.Redis.CmdTimeout)*time.Second)
	defer cancel()
	//
	return NewSession[T](ctx, prefixKey, data)
}

// GetSessionTimeout 获取会话
func GetSessionTimeout[T any](id string) (*Session[T], error) {
	// 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Cfg.Redis.CmdTimeout)*time.Second)
	defer cancel()
	//
	return GetSession[T](ctx, id)
}

// GetUserSession 获取用户会话
func GetUserSession(sessionID string) (*Session[*User], error) {
	return GetSessionTimeout[*User](sessionID)
}

// NewUserSession 创建用户会话
func NewUserSession(user *User) (*Session[*User], error) {
	return NewSessionTimeout[*User](UserSessionPrefix, user)
}

// GetDeveloperSession 获取开发者会话
func GetDeveloperSession(sessionID string) (*Session[*Developer], error) {
	return GetSessionTimeout[*Developer](sessionID)
}

// NewDeveloperSession 创建用户会话
func NewDeveloperSession(developer *Developer) (*Session[*Developer], error) {
	return NewSessionTimeout[*Developer](DeveloperSessionPrefix, developer)
}
