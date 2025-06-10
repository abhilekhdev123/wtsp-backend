package servey

import (
	"time"
	"wtsp-backend/server/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Option struct {
	Name        string `bson:"name" json:"name"`
	Value       string `bson:"value" json:"value"`
	Icon        string `bson:"icon" json:"icon"`
	Description string `bson:"description" json:"description"`
}

type Validation struct {
	Min int `bson:"min" json:"min"`
	Max int `bson:"max" json:"max"`
}

type Survey struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title           string             `bson:"title" json:"title"`
	SubTitle        string             `bson:"subTitle" json:"subTitle"`
	MetaTitle       string             `bson:"metaTitle" json:"metaTitle"`
	Description     string             `bson:"description" json:"description"`
	Category        string             `bson:"category" json:"category"`
	MetaDescription string             `bson:"metaDescription" json:"metaDescription"`
	IsOptional      bool               `bson:"isOptional" json:"isOptional"`
	Options         []Option           `bson:"options" json:"options"`
	OptionsLbs      []Option           `bson:"optionslbs" json:"optionslbs"`
	IsMultiSelect   bool               `bson:"isMultiSelect" json:"isMultiSelect"`
	SeqNo           int                `bson:"seqNo" json:"seqNo"`
	Images          []string           `bson:"images" json:"images"`
	Validation      Validation         `bson:"validation" json:"validation"`
	ProfileKey      string             `bson:"profileKey" json:"profileKey"`
	Exclusion       bool               `bson:"exclusion" json:"exclusion"`
	CreatedBy       primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy"`
	Deleted         bool               `bson:"deleted" json:"deleted"`

	// Timestamps
	CreatedAt time.Time `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty" json:"updatedAt"`

	// mongoose-delete plugin fields
	DeletedAt *time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	DeletedBy *primitive.ObjectID `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
}

var SurveyCollection *mongo.Collection

func Init() error {
	SurveyCollection = config.MongoDB.Collection("Survey")
	return nil
}
