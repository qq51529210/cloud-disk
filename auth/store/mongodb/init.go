package mongodb

import (
	"context"
	"time"

	"github.com/qq51529210/micro-services/auth/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseName = "micro-service-auth"
	queryTimeout = time.Second * 3
)

func Init(cfg map[string]interface{}) store.Store {
	if v, ok := cfg["databaseName"].(string); ok {
		databaseName = v
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
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
