package db

import (
	"context"

	"github.com/tripathysagar/transport-booking/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	PostUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoStore(c *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client: c,
		coll:   c.Database(dbname).Collection(userColl),
	}
}

func (ms *MongoUserStore) PostUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := ms.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}
