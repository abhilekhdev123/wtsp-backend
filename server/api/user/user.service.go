package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"wtsp-backend/server/api/auth"
	"wtsp-backend/server/utility"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func CreateUserService(req *CreateUserRequest) (*User, int, string, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"email": req.Email},
			{"phone": req.Phone},
		},
	}

	var existing User
	err := userCollection.FindOne(context.TODO(), filter).Decode(&existing)
	if err == nil {
		return nil, http.StatusConflict, "User with this email or phone already exists", nil
	}

	user := User{
		ID:                 primitive.NewObjectID(),
		Email:              req.Email,
		Name:               req.Name,
		Phone:              req.Phone,
		Role:               req.Role,
		ProfilePic:         req.ProfilePic,
		TargetPlatform:     req.TargetPlatform,
		DeviceToken:        req.DeviceToken,
		IsGoogleLogin:      getBool(req.IsGoogleLogin),
		IsEmailVerified:    false,
		IsMobileVerified:   false,
		IsActive:           true,
		NotificationStatus: true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	userSalt := utility.GenerateSalt() // generate a random salt
	user.HashSalt = userSalt
	// Hash the password if provided
	password := req.Password
	if password == "" {
		password = "123456"
	}
	user.HashedPassword = utility.HashPassword(password, userSalt)

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, http.StatusInternalServerError, "Failed to create user", err
	}

	// token, refreshtoken, err := auth.GenerateJWT(&user)
	token, refreshtoken, err := auth.GenerateJWT(user.ID.Hex(), user.Role)

	if err != nil {
		return nil, http.StatusInternalServerError, "Failed to generate JWT", err
	}

	fmt.Println("token", token, "refreshtoken", refreshtoken)

	return &user, http.StatusOK, "User created successfully", nil
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
