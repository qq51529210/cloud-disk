package redis

import (
	"errors"
	"fmt"

	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/redis"
)

var (
	ErrNoInteger = errors.New("is not a integer")
)

func cfgError(key string, value interface{}) error {
	return fmt.Errorf("config.redis.%s invalid value <%v>", key, value)
}

func New(cfg map[string]interface{}) cache.Cache {
	ca := new(Cache)
	rc := redis.ClientConfig{}
	rc.Load(cfg)
	ca.Client = redis.NewClient(&rc)
	return ca
}

type Cache struct {
	*redis.Client
}

func (ca *Cache) Set(key string, value interface{}, expired int) error {
	_, err := ca.Client.Cmd("SET", key, value, "EX", expired)
	return err
}

func (ca *Cache) Get(key string) (interface{}, error) {
	return ca.Client.Cmd("GET", key)
}

func (ca *Cache) Has(key string) (bool, error) {
	v, err := ca.Client.Cmd("EXISTS", key)
	if err != nil {
		return false, err
	}
	if num, ok := v.(int64); ok {
		return num == 1, nil
	}
	return false, ErrNoInteger
}

func (ca *Cache) Del(key string) error {
	_, err := ca.Client.Cmd("DEL", key)
	return err
}
