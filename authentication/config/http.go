package config

import (
	"errors"
	"os"
)

// httpConfig 表示 web 服务的配置
type httpConfig struct {
	// 监听地址
	Address string `json:"address" yaml:"address" xml:"Address"`
	// X509 公钥
	CertPEM string `json:"certPEM" yaml:"certPEM" xml:"CertPEM"`
	// X509 密钥
	KeyPEM string `json:"keyPEM" yaml:"keyPEM" xml:"KeyPEM"`
}

// Check 检测字段
func (c *httpConfig) Check() {
	// Address
	if c.Address == "" {
		panic(errors.New("empty http address"))
	}
}

// ReadEnv 使用环境变量替换掉配置的字段
func (c *httpConfig) ReadEnv() {
	// Address
	str := os.Getenv("HTTP_ADDRESS")
	if str != "" {
		c.Address = str
	}
}
