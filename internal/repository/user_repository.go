package repository

import (
	"context"
	"go_youtube_at_home/configs"
	"go_youtube_at_home/internal/domain"
	databaseModel "go_youtube_at_home/internal/model/database_model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection 		*mongo.Collection
	mongoTimeout 	int
	// ctx 				context.Context
}

func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &userRepository{
		collection: collection,
		mongoTimeout: configs.GetConfig().MongoDB.Timeout,
		// ctx,
	}
}

func (r *userRepository) CreateUser(user *databaseModel.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.mongoTimeout)*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) GetUserByUsername(username string) (*databaseModel.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.mongoTimeout)*time.Second)
	defer cancel()
	
	var user databaseModel.User
	err := r.collection.FindOne(ctx, bson.D{{"username", username}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// open /storage/862efcd4-442d-4b8b-b290-90fca8fce126: no such file or directory