package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Email              string              `bson:"email" json:"email"` // required
	Name               string              `bson:"name,omitempty" json:"name,omitempty"`
	Phone              string              `bson:"phone,omitempty" json:"phone,omitempty"`
	Role               []string            `bson:"role" json:"role"` // required
	HashedPassword     string              `bson:"hashedPassword,omitempty" json:"-"`
	PasswordExpiry     *time.Time          `bson:"passowrdExpiry,omitempty" json:"passowrdExpiry,omitempty"`
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
	IsPrime            bool                `bson:"isPrime" json:"isPrime"`
	NotificationStatus bool                `bson:"notificationstatus" json:"notificationstatus"`
	IsGoogleLogin      bool                `bson:"isGoogleLogin,omitempty" json:"isGoogleLogin"`
	CreatedAt          time.Time           `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time           `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	DeletedAt          *time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	DeletedBy          *primitive.ObjectID `bson:"deletedBy,omitempty" json:"deletedBy,omitempty"`
}
