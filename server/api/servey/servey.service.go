package servey

import (
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateServeyService(req *CreateServeyRequest) (*Survey, int, string, error) {
	servey := Survey{
		ID:              primitive.NewObjectID(),
		Title:           req.Title,
		SubTitle:        req.SubTitle,
		MetaTitle:       req.MetaTitle,
		Description:     req.Description,
		Category:        req.Category,
		MetaDescription: req.MetaDescription,
		IsOptional:      req.IsOptional,
		Options:         req.Options,
		OptionsLbs:      req.OptionsLbs,
		IsMultiSelect:   req.IsMultiSelect,
		SeqNo:           req.SeqNo,
		Images:          req.Images,
		Validation:      req.Validation,
		ProfileKey:      req.ProfileKey,
		Exclusion:       req.Exclusion,
		Deleted:         req.Deleted,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	_, err := SurveyCollection.InsertOne(context.TODO(), servey)
	if err != nil {
		return nil, http.StatusInternalServerError, "Failed to create survey", err
	}

	return &servey, http.StatusOK, "Survey created successfully", nil
}

func GetServeyListFromService() ([]Survey, error) {
	var serveys []Survey

	//////i want to get all users, so i will not use any filter
	cursor, err := SurveyCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err // Error finding documents
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var servey Survey
		if err := cursor.Decode(&servey); err != nil {
			return nil, err
		}
		serveys = append(serveys, servey)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return serveys, nil
}
