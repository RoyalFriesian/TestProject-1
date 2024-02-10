package utils

import (
	"context"

	"github.com/RoyalFriesian/TestProject-1/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveUserInMongo(user models.User) error {
	clientOptions := options.Client().ApplyURI(MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	collection := client.Database("test").Collection("users")
	_, err = collection.InsertOne(context.TODO(), user)
	return err
}
