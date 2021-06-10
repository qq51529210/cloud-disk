package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qq51529210/redis"
)

var (
	tokenRedis *redis.Client
	phoneRedis *redis.Client
	mysqlDB    *sql.DB
)

func InitRedis(tokenUrl, phoneUrl string) error {
	var err error
	if tokenUrl == "" {
		tokenUrl = "redis://127.0.0.1:6379?db=1&max_conn=10&read_time=1000&write_time=1000"
	}
	tokenRedis, err = redis.NewClient(nil, tokenUrl)
	if err != nil {
		return err
	}
	if phoneUrl == "" {
		phoneUrl = "redis://127.0.0.1:6379?db=2&max_conn=10&read_time=1000&write_time=1000"
	}
	phoneRedis, err = redis.NewClient(nil, phoneUrl)
	if err != nil {
		return err
	}
	return nil
}

func InitMysql(url string) error {
	var err error
	if url == "" {
		url = "root:123456@tcp(127.0.0.1:3306)/anthentication"
	}
	mysqlDB, err = sql.Open("mysql", url)
	if err != nil {
		return err
	}
	return mysqlDB.Ping()
}
