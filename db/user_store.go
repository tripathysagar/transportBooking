package db

import (
	"context"

	"github.com/tripathysagar/transport-booking/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	GetUsers(context.Context) ([]*types.User, error)
	GetUserByID(context.Context, string) (*types.User, error)
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

func (ms *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	res, err := ms.coll.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []*types.User
	if err := res.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
func (ms *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := ms.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ms *MongoUserStore) PostUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := ms.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}
