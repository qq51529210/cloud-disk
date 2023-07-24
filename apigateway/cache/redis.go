package cache

import (
	"apigateway/cfg"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rds redis.UniversalClient
)

// Init 初始化缓存
func Init() error {
	rds = redis.NewUniversalClient(&redis.UniversalOptions{
		ClientName:       cfg.Cfg.Redis.Name,
		Addrs:            cfg.Cfg.Redis.Addrs,
		DB:               cfg.Cfg.Redis.DB,
		Username:         cfg.Cfg.Redis.Username,
		Password:         cfg.Cfg.Redis.Password,
		MasterName:       cfg.Cfg.Redis.Master,
		SentinelUsername: cfg.Cfg.Redis.SentinelUsername,
		SentinelPassword: cfg.Cfg.Redis.SentinelPassword,
	})
	return rds.Ping(context.Background()).Err()
}

func newRedisTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(cfg.Cfg.Redis.CmdTimeout)*time.Second)
}

// PutWithContext 保存缓存
func PutWithContext(ctx context.Context, k string, v any, t int64) error {
	data, _ := json.Marshal(v)
	if t > 0 {
		return rds.Set(ctx, k, data, time.Duration(t)*time.Second).Err()
	}
	return rds.Set(ctx, k, data, -1).Err()
}

// Put 保存缓存
func Put(k string, v any, t int64) error {
	ctx, cancel := newRedisTimeout()
	defer cancel()
	return PutWithContext(ctx, k, v, t)
}

// GetWithContext 查询缓存
func GetWithContext[T any](ctx context.Context, k string) (*T, error) {
	data, err := rds.Get(ctx, k).Bytes()
	if err != nil {
		if err == redis.Nil {
			err = nil
		}
		return nil, err
	}
	//
	t := new(T)
	_ = json.Unmarshal(data, t)
	//
	return t, nil
}

// Get 查询缓存
func Get[T any](k string) (*T, error) {
	ctx, cancel := newRedisTimeout()
	defer cancel()
	return GetWithContext[T](ctx, k)
}

// DelWithContext 删除缓存
func DelWithContext(ctx context.Context, k string) error {
	return rds.Del(ctx, k).Err()
}

// Del 删除缓存
func Del(k string) error {
	ctx, cancel := newRedisTimeout()
	defer cancel()
	return rds.Del(ctx, k).Err()
}
