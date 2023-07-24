package cache

// Server 表示一个服务器
type Server struct {
	ID string
	// 服务
	BaseURL string
	// 负载
	Load int64
}

func AddServerLoad(path,server) {
	rds.ZIncrBy()
}