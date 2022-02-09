package redis

import (
	"github.com/qq51529210/micro-services/auth/cache"
	"github.com/qq51529210/redis"
	"github.com/qq51529210/uuid"
	"github.com/qq51529210/web/util"
)

var (
	tokenDatabase     = 0
	phoneCodeDatabase = 1
	tokenExpire       = 60 * 10
	phoneCodeLength   = 6
	phoneCodeExpire   = 60
)

func Init(cfg map[string]interface{}) cache.Cache {
	if v, ok := cfg["tokenDatabase"].(float64); ok {
		tokenDatabase = int(v)
	}
	if v, ok := cfg["phoneCodeDatabase"].(float64); ok {
		phoneCodeDatabase = int(v)
	}
	if v, ok := cfg["tokenExpire"].(float64); ok {
		tokenExpire = int(v)
	}
	if v, ok := cfg["phoneCodeLength"].(float64); ok {
		phoneCodeLength = int(v)
	}
	if v, ok := cfg["phoneCodeExpire"].(float64); ok {
		phoneCodeExpire = int(v)
	}
	//
	c := new(Cache)
	cc := redis.ClientConfig{}
	cc.Load(cfg)
	c.client = redis.NewClient(&cc)
	return c
}

type Cache struct {
	client *redis.Client
}

func (c *Cache) NewToken(value string) (string, error) {
	token := uuid.LowerV1WithoutHyphen()
	_, err := c.client.Cmd("SET", token, value, "EX", tokenExpire)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Cache) SetToken(token, value string) error {
	_, err := c.cmd(tokenDatabase, "SET", token, value, "EX", tokenExpire)
	return err
}

func (c *Cache) HasToken(token string) (bool, error) {
	res, err := c.cmd(tokenDatabase, "EXISTS", token)
	if err != nil {
		return false, err
	}
	if v, ok := res.(int64); ok {
		return v == 1, nil
	}
	return false, redis.Error("invalid response data type")
}

func (c *Cache) GetToken(token string) (string, error) {
	res, err := c.cmd(tokenDatabase, "GET", token)
	if err != nil {
		return "", err
	}
	if v, ok := res.(string); ok {
		return v, nil
	}
	return "", redis.Error("invalid response data type")
}

func (c *Cache) UpdateToken(token string) error {
	res, err := c.cmd(tokenDatabase, "EXPIRE", token)
	if err != nil {
		return err
	}
	if v, ok := res.(int64); ok {
		if v == 1 {
			return nil
		}
		return redis.Error("update fail")
	}
	return redis.Error("invalid response data type")
}

func (c *Cache) DeleteToken(token string) error {
	res, err := c.cmd(tokenDatabase, "DEL", token)
	if err != nil {
		return err
	}
	if v, ok := res.(int64); ok {
		if v == 1 {
			return nil
		}
		return redis.Error("delete fail")
	}
	return redis.Error("invalid response data type")
}

func (c *Cache) NewPhoneCode(number string) (string, error) {
	code := util.RandomNumber(phoneCodeLength)
	_, err := c.cmd(phoneCodeDatabase, "SET", number, code, "EX", phoneCodeExpire)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (c *Cache) GetPhoneCode(number string) (string, error) {
	res, err := c.cmd(phoneCodeDatabase, "GET", number)
	if err != nil {
		return "", err
	}
	//
	if v, ok := res.(string); ok {
		return v, nil
	}
	return "", redis.Error("invalid response data type")
}

func (c *Cache) cmd(database int, args ...interface{}) (interface{}, error) {
	conn, err := c.client.Conn()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	//
	err = conn.Database(database)
	if err != nil {
		return "", err
	}
	return conn.Cmd(args...)
}
