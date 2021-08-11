package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/qq51529210/redis"
)

var (
	session *redis.Client
	code    *redis.Client
	mysqlDB *sql.DB
)

type Config struct {
	Session  string `json:"session"`
	VeriCode string `json:"verificationCode"`
	Mysql    string `json:"mysq"`
}

func Init(c *Config) {
	var err error
	// Session redis
	if c.Session == "" {
		session, err = redis.NewClient(nil, "redis://127.0.0.1:6379?db=0&max_conn=10&read_time=1000&write_time=1000")
	} else {
		session, err = redis.NewClient(nil, c.Session)
	}
	if err != nil {
		panic(err)
	}
	// Verification Code redis
	if c.VeriCode == "" {
		code, err = redis.NewClient(nil, "redis://127.0.0.1:6379?db=1&max_conn=10&read_time=1000&write_time=1000")
	} else {
		code, err = redis.NewClient(nil, c.Session)
	}
	if err != nil {
		panic(err)
	}
	// Mysql
	if c.Mysql == "" {
		mysqlDB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/auth")
	} else {
		mysqlDB, err = sql.Open("mysql", c.Mysql)
	}
	if err != nil {
		panic(err)
	}
	initUserStmt(mysqlDB)
}

func IsExistedError(err error) bool {
	if e, ok := err.(*mysql.MySQLError); ok {
		return e.Number == 1062
	}
	return false
}
