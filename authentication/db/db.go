package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qq51529210/redis"
)

var (
	tokenRedis *redis.Client
	phoneRedis *redis.Client
	mysqlDB    *sql.DB
)

func InitTokenRedis(url string) error {
	if url == "" {
		url = "redis://127.0.0.1:6379?db=1&max_conn=10&read_time=1000&write_time=1000"
	}
	rds, err := redis.NewClient(nil, url)
	if err != nil {
		return err
	}
	if tokenRedis != nil {
		tokenRedis.Close()
	}
	tokenRedis = rds
	return nil
}

func InitPhoneNumberRedis(url string) error {
	if url == "" {
		url = "redis://127.0.0.1:6379?db=1&max_conn=20&read_time=1000&write_time=1000"
	}
	rds, err := redis.NewClient(nil, url)
	if err != nil {
		return err
	}
	if phoneRedis != nil {
		phoneRedis.Close()
	}
	phoneRedis = rds
	return nil
}

func InitMysql(url string) error {
	if url == "" {
		url = "root:123456@tcp(127.0.0.1:3306)/anthentication"
	}
	db, err := sql.Open("mysql", url)
	if err != nil {
		return err
	}
	err = initUserStmt(db)
	if err != nil {
		return err
	}
	if mysqlDB != nil {
		mysqlDB.Close()
	}
	mysqlDB = db
	return nil
}

func IsExistedError(err error) bool {
	if e, ok := err.(*mysql.MySQLError); ok {
		return e.Number == 1062
	}
	return false
}
