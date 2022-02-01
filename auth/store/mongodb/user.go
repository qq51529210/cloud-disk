package mongodb

import (
	"context"

	"github.com/qq51529210/micro-services/auth/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Store) AddUser(model *store.UserModel) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	//
	res, err := s.userCollection.InsertOne(ctx, bson.M{
		"account":  model.Account,
		"password": model.Password,
		"phone":    model.Phone,
	})
	if err != nil {
		return "", err
	}
	_id := res.InsertedID.(primitive.ObjectID)
	return string(_id[0:]), nil
}

func (s *Store) DeleteUser(_id string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	//
	res, err := s.userCollection.DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (s *Store) UpdateUserPassword(model *store.UserModel) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	//
	res, err := s.userCollection.UpdateByID(ctx, bson.M{"_id": model.ID}, bson.M{"password": model.Password})
	if err != nil {
		return 0, err
	}
	return res.UpsertedCount, nil
}

func (s *Store) GetUser(account string) (*store.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	//
	res := s.userCollection.FindOne(ctx, bson.M{
		"$or": bson.A{
			bson.E{Key: "account", Value: account},
			bson.E{Key: "phone", Value: account},
		},
	})
	var model store.UserModel
	err := res.Decode(&model)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (s *Store) GetUserList(query *store.PageQueryModel) (*store.PageDataModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
	defer cancel()
	//
	var err error
	var page store.PageDataModel
	page.Count, err = s.userCollection.EstimatedDocumentCount(ctx, &options.EstimatedDocumentCountOptions{
		MaxTime: &queryTimeout,
	})
	if err != nil {
		return nil, err
	}
	cur, err := s.userCollection.Find(ctx, bson.D{}, &options.FindOptions{
		MaxTime: &queryTimeout,
		Projection: bson.A{
			bson.E{Key: "_id", Value: 1},
			bson.E{Key: "account", Value: 1},
			bson.E{Key: "phone", Value: 1},
		},
		Skip:  &query.Offset,
		Limit: &query.Count,
	})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var model store.UserModel
		err = cur.Decode(&model)
		if err != nil {
			return nil, err
		}
		page.Data = append(page.Data, &model)
	}
	return &page, nil
}
