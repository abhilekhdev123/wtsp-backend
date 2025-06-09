// package user

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type User struct {
// 	ID                 primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
// 	Email              string              `bson:"email" json:"email"` // required
// 	Name               string              `bson:"name,omitempty" json:"name,omitempty"`
// 	Phone              string              `bson:"phone,omitempty" json:"phone,omitempty"`
// 	Role               []string            `bson:"role" json:"role"` // required
// 	HashedPassword     string              `bson:"hashedPassword,omitempty" json:"-"`
// 	PasswordExpiry     *time.Time          `bson:"passowrdExpiry,omitempty" json:"passowrdExpiry,omitempty"`
// 	HashSalt           string              `bson:"hashSalt,omitempty" json:"-"`
// 	IsEmailVerified    bool                `bson:"isEmailVerified" json:"isEmailVerified"`
// 	IsMobileVerified   bool                `bson:"isMobileVerified" json:"isMobileVerified"`
// 	IsActive           bool                `bson:"isActive" json:"isActive"`
// 	ProfilePic         string              `bson:"profilePic,omitempty" json:"profilePic,omitempty"`
// 	ResetPasswordToken string              `bson:"resetPasswordToken,omitempty" json:"-"`
// 	CreatedBy          *primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
// 	DeviceID           string              `bson:"deviceId,omitempty" json:"deviceId,omitempty"`
// 	TargetPlatform     string              `bson:"targetPlatform,omitempty" json:"targetPlatform,omitempty"`
// 	DeviceToken        string              `bson:"deviceToken,omitempty" json:"deviceToken,omitempty"`
// 	IPAddress          string              `bson:"ipaddress,omitempty" json:"ipaddress,omitempty"`
// 	IsPrime            bool                `bson:"isPrime" json:"isPrime"`
// 	NotificationStatus bool                `bson:"notificationstatus" json:"notificationstatus"`
// 	IsGoogleLogin      bool                `bson:"isGoogleLogin,omitempty" json:"isGoogleLogin"`
// 	CreatedAt          time.Time           `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
// 	UpdatedAt          time.Time           `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
// 	DeletedAt          *time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
// 	DeletedBy          *primitive.ObjectID `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
// }

package user

import (
	"context"
	"errors"
	"time"
	"wtsp-backend/server/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ValidRoles defines allowed user roles
var ValidRoles = []string{"Admin", "User", "Manager", "Coach"}

// User represents the user document in MongoDB
type User struct {
	ID                 primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Email              string              `bson:"email" json:"email"` // Unique + Indexed
	Name               string              `bson:"name,omitempty" json:"name,omitempty"`
	Phone              string              `bson:"phone,omitempty" json:"phone,omitempty"` // Unique + Indexed
	Role               []string            `bson:"role" json:"role"`                       // Should be validated against ValidRoles
	HashedPassword     string              `bson:"hashedPassword,omitempty" json:"-"`
	PasswordExpiry     *time.Time          `bson:"passwordExpiry,omitempty" json:"passwordExpiry,omitempty"`
	HashSalt           string              `bson:"hashSalt,omitempty" json:"-"`
	IsEmailVerified    bool                `bson:"isEmailVerified" json:"isEmailVerified"`
	IsMobileVerified   bool                `bson:"isMobileVerified" json:"isMobileVerified"`
	IsActive           bool                `bson:"isActive" json:"isActive"`
	ProfilePic         string              `bson:"profilePic,omitempty" json:"profilePic,omitempty"`
	ResetPasswordToken string              `bson:"resetPasswordToken,omitempty" json:"-"`
	CreatedBy          *primitive.ObjectID `bson:"createdBy,omitempty" json:"createdBy,omitempty"`
	DeviceID           string              `bson:"deviceId,omitempty" json:"deviceId,omitempty"`
	TargetPlatform     string              `bson:"targetPlatform,omitempty" json:"targetPlatform,omitempty"`
	DeviceToken        string              `bson:"deviceToken,omitempty" json:"deviceToken,omitempty"`
	IPAddress          string              `bson:"ipaddress,omitempty" json:"ipaddress,omitempty"`
	IsPrime            bool                `bson:"isPrime" json:"isPrime"` // default false
	NotificationStatus bool                `bson:"notificationstatus" json:"notificationstatus"`
	IsGoogleLogin      bool                `bson:"isGoogleLogin,omitempty" json:"isGoogleLogin"`
	IsAppleLogin       bool                `bson:"isAppleLogin,omitempty" json:"isAppleLogin"`
	AppleUserID        string              `bson:"appleUserId,omitempty" json:"appleUserId,omitempty"`
	InitialWeight      string              `bson:"initialWeight,omitempty" json:"initialWeight,omitempty"` // set default manually
	BMI                float64             `bson:"bmi,omitempty" json:"bmi,omitempty"`
	BMR                string              `bson:"bmr,omitempty" json:"bmr,omitempty"`
	CreatedAt          time.Time           `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time           `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt          *time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	DeletedBy          *primitive.ObjectID `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
}

// ValidateRoles checks that each role in user roles exists in ValidRoles
func ValidateRoles(roles []string) error {
	for _, role := range roles {
		if !isValidRole(role) {
			return errors.New("invalid role: " + role)
		}
	}
	return nil
}

// isValidRole helper
func isValidRole(role string) bool {
	for _, valid := range ValidRoles {
		if role == valid {
			return true
		}
	}
	return false
}

// CreateIndexes creates unique indexes for email and phone fields on the user collection
func CreateIndexes(ctx context.Context, userCollection *mongo.Collection) error {
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("unique_email"),
		},
		{
			Keys:    bson.D{{Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true).SetName("unique_phone"),
		},
	}

	_, err := userCollection.Indexes().CreateMany(ctx, indexModels)
	return err
}

func Init() error {
	userCollection = config.MongoDB.Collection("users")
	ctx := context.Background()
	return CreateIndexes(ctx, userCollection)
}
