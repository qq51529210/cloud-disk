package db

import (
	"context"
	"oauth2/cfg"

	"github.com/qq51529210/util"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	_db *gorm.DB
	rds redis.UniversalClient
)

// 用于取地址
var (
	True  int8 = 1
	False int8 = 0
)

// Init 初始化数据库
func Init() error {
	var err error
	// 初始化
	dbCfg := util.NewGORMConfig()
	dbCfg.Logger = &util.GORMLog{}
	_db, err = util.InitGORM(cfg.Cfg.DB.URL, dbCfg)
	if err != nil {
		return err
	}
	// 数据表
	err = initTable()
	if err != nil {
		return err
	}
	// 缓存
	err = initReids()
	if err != nil {
		return err
	}
	//
	return nil
}

// initTable 初始化数据表
func initTable() error {
	_db.AutoMigrate(
		new(User),
		new(Developer),
		new(Client),
	)
	//
	return nil
}

// initReids 初始化缓存
func initReids() error {
	rds = redis.NewUniversalClient(&redis.UniversalOptions{
		ClientName:       cfg.Cfg.Redis.Name,
		Addrs:            cfg.Cfg.Redis.Addrs,
		DB:               cfg.Cfg.Redis.DB,
		Username:         cfg.Cfg.Redis.Username,
		Password:         cfg.Cfg.Redis.Password,
		MasterName:       cfg.Cfg.Redis.Master,
		SentinelUsername: cfg.Cfg.Redis.SentinelUsername,
		SentinelPassword: cfg.Cfg.Redis.SentinelPassword,
	})
	return rds.Ping(context.Background()).Err()
}
