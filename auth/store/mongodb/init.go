package mongodb

import (
	"context"
	"time"

	"github.com/qq51529210/micro-services/auth/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseName = "auth"
	queryTimeout = time.Second * 3
)

func Init(cfg map[string]interface{}) store.Store {
	var uri string
	if v, ok := cfg["uri"].(string); ok {
		uri = v
	} else {
		uri = "mongodb://localhost:27017"
	}
	if v, ok := cfg["databaseName"].(string); ok {
		databaseName = v
	}
	if v, ok := cfg["queryTimeout"].(float64); ok {
		n := time.Duration(v) * time.Millisecond
		if n < 1 {
			queryTimeout = time.Second * 3
		} else {
			queryTimeout = n
		}
	}
	opt := options.Client()
	opt.ApplyURI(uri)
	opt.SetAppName(databaseName)
	cred := options.Credential{}
	if v, ok := cfg["username"].(string); ok {
		cred.Username = v
	}
	if v, ok := cfg["password"].(string); ok {
		cred.Password = v
	}
	opt.SetAuth(cred)
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	s := new(Store)
	s.client = client
	database := s.client.Database(databaseName)
	s.userCollection = database.Collection("user")
	return s
}

type Store struct {
	client         *mongo.Client
	userCollection *mongo.Collection
}
