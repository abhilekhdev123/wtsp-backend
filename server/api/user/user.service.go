package user

import (
	"context"
	"time"
	"wtsp-backend/server/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = config.MongoDB.Collection("users")
}

func CreateUserService(u User) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	_, err := userCollection.InsertOne(context.TODO(), u)
	return err
}
