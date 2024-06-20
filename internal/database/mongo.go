package database

import (
	"context"
	"go_youtube_at_home/configs"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var videoCollection *mongo.Collection

func InitMongo() error {

	connectionString := configs.GetConfig().MongoDB.URI
	databaseName := configs.GetConfig().MongoDB.DatabaseName

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	userCollection = client.Database(databaseName).Collection("user")
	videoCollection = client.Database(databaseName).Collection("video")

	return nil
}

func GetUserCollection() *mongo.Collection {
	return userCollection
}

func GetVideoCollection() *mongo.Collection {
	return videoCollection
}

