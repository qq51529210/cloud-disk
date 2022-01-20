package redis

import (
	"errors"
	"fmt"
	"time"

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
	ca := new(_Cache)
	var host string
	var cap int
	var rto time.Duration
	var wto time.Duration
	v, ok := cfg["host"]
	if !ok {
		s, ok := v.(string)
		if !ok {
			panic(cfgError("host", v))
		}
		host = s
	}
	v, ok = cfg["maxConn"]
	if ok {
		s, ok := v.(float64)
		if !ok {
			panic(cfgError("maxConn", v))
		}
		cap = int(s)
	}
	v, ok = cfg["readTimeout"]
	if ok {
		s, ok := v.(float64)
		if !ok {
			panic(cfgError("readTimeout", v))
		}
		rto = time.Duration(s) * time.Millisecond
	}
	v, ok = cfg["writeTimeout"]
	if ok {
		s, ok := v.(float64)
		if !ok {
			panic(cfgError("writeTimeout", v))
		}
		wto = time.Duration(s) * time.Millisecond
	}
	ca.Client = redis.NewClient(redis.NewPool(host, cap, rto, wto))
	return ca
}

type _Cache struct {
	*redis.Client
}

func (ca *_Cache) Set(key string, value interface{}, expired int) error {
	_, err := ca.Client.Command("SET", key, value, "EX", expired)
	return err
}

func (ca *_Cache) Get(key string) (interface{}, error) {
	return ca.Client.Command("GET", key)
}

func (ca *_Cache) Has(key string) (bool, error) {
	v, err := ca.Client.Command("EXISTS", key)
	if err != nil {
		return false, err
	}
	if num, ok := v.(int64); ok {
		return num == 1, nil
	}
	return false, ErrNoInteger
}

func (ca *_Cache) Del(key string) error {
	_, err := ca.Client.Command("DEL", key)
	return err
}
