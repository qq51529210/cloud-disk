package cfg

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/qq51529210/util"
)

var (
	// Cfg 实例
	Cfg Config
)

// Config 表示程序启动配置
type Config struct {
	// 服务名称，日志使用
	Name string `json:"name" yaml:"name" validate:"required,max=32"`
	// 监听地址
	Addr string `json:"addr" yaml:"addr" validate:"required"`
	// 测试地址
	Test string `json:"test" yaml:"test"`
	// 日志配置
	Log util.LogCfg `json:"log" yaml:"log"`
	// 数据库配置
	DB struct {
		URL string `json:"url" yaml:"url" validate:"required"`
	} `json:"db" yaml:"db"`
	// 缓存配置
	Redis struct {
		// 客户端名称
		Name string `json:"name" yaml:"name"`
		// 服务地址
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required"`
		// 数据库
		DB int `json:"db" yaml:"db" validate:"min=0,max=32"`
		// 用户名 6.0 以上
		Username string `json:"username" yaml:"username"`
		// 密码
		Password string `json:"password" yaml:"password"`
		// 集群
		Master string `json:"master" yaml:"master"`
		// 哨兵用户名
		SentinelUsername string `json:"sentinelUsername" yaml:"sentinelUsername"`
		// 哨兵密码
		SentinelPassword string `json:"sentinelPassword" yaml:"sentinelPassword"`
		// 执行一次命令的超时时间
		CmdTimeout int64 `json:"cmdTimeout" yaml:"cmdTimeout" validate:"required,min=1"`
	} `json:"redis" yaml:"redis"`
	// 会话配置
	Session struct {
		// 过期时间
		Expires int64 `json:"expires" yaml:"expires" validate:"required,min=60"`
	} `json:"session" yaml:"session"`
	// oauth2
	OAuth2 struct {
		// 是否启用 Implicit 授权模式
		EnableImplicitGrant bool `json:"enableImplicitGrant" yaml:"enableImplicitGrant"`
		// 是否启用 Password 授权模式
		EnablePasswordGrant bool `json:"enablePasswordGrant" yaml:"enablePasswordGrant"`
		// 是否启用 Client 授权模式
		EnableClientGrant bool `json:"enableClientGrant" yaml:"enableClientGrant"`
		// 授权码过期时间
		AuthorizationCodeExpires int64 `json:"authorizationCodeExpires" yaml:"authorizationCodeExpires" validate:"required,min=3"`
		// 访问令牌过期时间
		AccessTokenExpires int64 `json:"accessTokenExpires" yaml:"accessTokenExpires" validate:"required,min=60"`
		// 刷新令牌过期时间
		RefreshTokenExpires int64 `json:"refreshTokenExpires" yaml:"refreshTokenExpires" validate:"required,min=60"`
		// 访问令牌的类型
		AccessTokenType string `json:"accessTokenType" yaml:"accessTokenType" validate:"required,oneof=Bearer"`
	} `json:"oauth2" yaml:"oauth2"`
}

// Load 加载配置
func Load() error {
	// 路径，默认程序同目录下的 cfg.yaml 文件
	path := "cfg.yaml"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	// 加载
	err := util.ReadCfg(path, &Cfg)
	if err != nil {
		return err
	}
	// 检查字段
	val := validator.New()
	err = val.Struct(&Cfg)
	if err != nil {
		return err
	}
	//
	return nil
}
