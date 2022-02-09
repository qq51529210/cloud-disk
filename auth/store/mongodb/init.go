package mongodb

import (
	"context"
	"time"

	"github.com/qq51529210/micro-services/auth/store"
	"go.mongodb.org/mongo-driver/bson"
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
	username := cfg["username"].(string)
	password := cfg["password"].(string)
	if username != "" && password != "" {
		opt.SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
	}
	client, err := mongo.Connect(context.Background(), opt)
	if err != nil {
		panic(err)
	}
	s := new(Store)
	s.client = client
	database := s.client.Database(databaseName)
	s.userCollection = database.Collection("user")
	idx := s.userCollection.Indexes()
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	_, err = idx.CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{bson.E{Key: "account", Value: 1}},
	})
	if err != nil {
		panic(err)
	}
	_, err = idx.CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{bson.E{Key: "phone", Value: 1}},
	})
	if err != nil {
		panic(err)
	}
	return s
}

type Store struct {
	client         *mongo.Client
	userCollection *mongo.Collection
}
