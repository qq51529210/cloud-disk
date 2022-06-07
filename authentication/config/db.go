package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// dbConfig 表示数据库的配置
type dbConfig struct {
	// 数据库连接字符串
	DSN string `json:"dsn" yaml:"dsn" xml:"DSN"`
	// 标准库的字段
	MaxIdleConns int `json:"maxIdleConns" yaml:"maxIdleConns" xml:"MaxIdleConns"`
	// 标准库的字段
	MaxOpenConns int `json:"maxOpenConns" yaml:"maxOpenConns" xml:"MaxOpenConns"`
	// 标准库的字段
	ConnMaxIdleTime int `json:"connMaxIdleTime" yaml:"connMaxIdleTime" xml:"ConnMaxIdleTime"`
	// 标准库的字段
	ConnMaxLifetime int `json:"connMaxLifetime" yaml:"connMaxLifetime" xml:"ConnMaxLifetime"`
}

// Check 检测字段
func (c *dbConfig) Check() {
	// DSN
	if c.DSN == "" {
		panic(errors.New("empty database dsn"))
	}
}

// ReadEnv 使用环境变量替换掉配置的字段
func (c *dbConfig) ReadEnv() {
	// DSN
	str := os.Getenv("DB_DSN")
	if str != "" {
		c.DSN = str
	}
	// MaxIdleConns
	str = os.Getenv("DB_MAX_IDLE_CONNS")
	if str != "" {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil || n < 0 {
			panic(fmt.Errorf("env DB_MAX_IDLE_CONNS invalid value %s", str))
		}
		c.MaxIdleConns = int(n)
	}
	// MaxOpenConns
	str = os.Getenv("DB_MAX_OPEN_CONNS")
	if str != "" {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil || n < 0 {
			panic(fmt.Errorf("env DB_MAX_OPEN_CONNS invalid value %s", str))
		}
		c.MaxOpenConns = int(n)
	}
	// ConnMaxIdleTime
	str = os.Getenv("DB_CONN_MAX_IDLE_TIME")
	if str != "" {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil || n < 0 {
			panic(fmt.Errorf("env DB_CONN_MAX_IDLE_TIME invalid value %s", str))
		}
		c.ConnMaxIdleTime = int(n)
	}
	// ConnMaxLifetime
	str = os.Getenv("DB_CONN_MAX_LIFE_TIME")
	if str != "" {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil || n < 0 {
			panic(fmt.Errorf("env DB_CONN_MAX_LIFE_TIME invalid value %s", str))
		}
		c.ConnMaxLifetime = int(n)
	}
}
