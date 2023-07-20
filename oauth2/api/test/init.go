package test

import (
	"oauth2/cfg"
	"oauth2/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qq51529210/util"
)

var (
	//
	client    = "test-Client"
	user      = "test-user"
	developer = "test-developer"
	pwd       = "123123"
	state     = "test-state"
	//
	host       = "http://localhost"
	oauth2Host = "http://localhost"
	//
	apiCallTimeout = time.Second * 5
)

// Init 初始化路由
func Init(g gin.IRouter) {
	// 先加入模拟数据
	testDeveloperData()
	testUserData()
	testClientData()
	//
	host += cfg.Cfg.Test
	oauth2Host += cfg.Cfg.Addr
	//
	g.GET("/", login)
	g.GET("/oauth2", oauth2)
}

func testDeveloperData() {
	m, err := db.GetDeveloper(developer)
	if err != nil {
		panic(err)
	}
	if m != nil {
		return
	}
	m = new(db.Developer)
	m.ID = developer
	m.Account = developer
	password := util.SHA1String(pwd)
	m.Password = &password
	m.Enable = &db.True
	_, err = db.AddDeveloper(m)
	if err != nil {
		panic(err)
	}
}

func testClientData() {
	m, err := db.GetClient(client)
	if err != nil {
		panic(err)
	}
	if m != nil {
		return
	}
	m = new(db.Client)
	m.ID = client
	m.Name = &client
	m.Secret = &pwd
	m.Enable = &db.True
	m.DeveloperID = developer
	_, err = db.AddClient(m)
	if err != nil {
		panic(err)
	}
}

func testUserData() {
	m, err := db.GetUser(user)
	if err != nil {
		panic(err)
	}
	if m != nil {
		return
	}
	m = new(db.User)
	m.ID = user
	m.Account = user
	password := util.SHA1String(pwd)
	m.Password = &password
	m.Enable = &db.True
	_, err = db.AddUser(m)
	if err != nil {
		panic(err)
	}
}
