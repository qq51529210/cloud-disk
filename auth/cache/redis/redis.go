package redis

import "github.com/qq51529210/micro-services/auth/cache"

func Init(cfg map[string]interface{}) cache.Cache {
	c := new(Cache)
	return c
}

type Cache struct {
}

func (c *Cache) NewToken(value string) (string, error) {
	return "", nil
}

func (c *Cache) SetToken(token, value string) error {
	return nil
}

func (c *Cache) HasToken(token string) (bool, error) {
	return false, nil
}

func (c *Cache) GetToken(token string) (string, error) {
	return "", nil
}

func (c *Cache) DelToken(token string) error {
	return nil
}

func (c *Cache) NewPhoneCode(number string) (string, error) {
	return "", nil
}

func (c *Cache) GetPhoneCode(number string) (string, error) {
	return "", nil
}
