package cache

var (
	_Cache Cache
)

type Cache interface {
	NewToken(value string) (string, error)
	SetToken(token, value string) error
	HasToken(token string) (bool, error)
	GetToken(token string) (string, error)
	DelToken(token string) error
	NewPhoneCode(number string) (string, error)
	GetPhoneCode(number string) (string, error)
}

func SetCache(cache Cache) {
	_Cache = cache
}

func GetCache() Cache {
	return _Cache
}
