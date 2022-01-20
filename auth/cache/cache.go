package cache

var (
	_Cache Cache
)

type Cache interface {
	Set(key string, value interface{}, expired int) error
	Get(key string) (interface{}, error)
	Has(key string) (bool, error)
	Del(key string) error
}

func SetCache(ca Cache) {
	_Cache = ca
}

func GetCache() Cache {
	return _Cache
}

func Set(key string, value interface{}, expired int) error {
	return _Cache.Set(key, value, expired)
}

func Get(key string) (interface{}, error) {
	return _Cache.Get(key)
}

func Has(key string) (bool, error) {
	return _Cache.Has(key)
}

func Del(key string) error {
	return _Cache.Del(key)
}
