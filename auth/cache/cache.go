package cache

import (
	"fmt"

	"github.com/qq51529210/micro-services/auth/cache/redis"
)

var (
	_Token     Token
	_PhoneCode PhoneCode
)

func Init(_type string, data map[string]interface{}) {
	switch _type {
	case "", "redis":
		_Token, _PhoneCode = redis.Init()
	default:
		panic(fmt.Errorf("config.cache.type: unsupported cache <%s>", _type))
	}
}

type Token interface {
	New(value string) (string, error)
	Set(token, value string) error
	Has(token string) (bool, error)
	Get(token string) (string, error)
	Del(token string) error
}

type PhoneCode interface {
	New(number string) (string, error)
	Get(number string) (string, error)
}

func GetToken() Token {
	return _Token
}

func GetPhoneCode() PhoneCode {
	return _PhoneCode
}
