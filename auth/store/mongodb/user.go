package mongodb

import (
	"context"

	"github.com/qq51529210/micro-services/auth/store"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Store) GetUser(account string) (*store.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	//
	s.userCollection.FindOne(ctx, bson.M{})
	return nil, nil
}
