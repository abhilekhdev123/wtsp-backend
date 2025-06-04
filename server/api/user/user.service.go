package user

import (
	"context"
	"time"
	"wtsp-backend/server/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserCollection() {
	userCollection = config.MongoDB.Collection("users")
}

func CreateUserService1(u User) (*User, error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	result, err := userCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}

	// Fetch the inserted document using the inserted ID
	var createdUser User
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&createdUser)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func CreateUserService(u User) (*User, error) {
	u.ID = primitive.NewObjectID() // manually set the _id
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	_, err := userCollection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}

	// No need to fetch again – we already have all the data
	return &u, nil
}

func GetUserService() ([]User, error) {
	var users []User

	//////i want to get all users, so i will not use any filter
	cursor, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err // Error finding documents
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
