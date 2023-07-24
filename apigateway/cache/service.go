package cache

const (
	serviceKeyPrefix = "service:"
)

// Service 表示服务组的缓存
type Service struct {
	Server  string
}

// GetServiceByPath 查询服务
func GetServiceByPath(path string) (*Service, error) {
	return nil, nil
}
