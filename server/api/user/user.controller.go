package user

import (
	"net/http"
	"time"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	// Safe assert - already validated by middleware
	req := c.MustGet("validatedBody").(*CreateUserRequest)

	user := User{
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

	// You can hash password here if needed:
	// user.HashedPassword = hashPassword(req.Password)

	createdUser, err := CreateUserService(user)
	if err != nil {
		config.SendError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	config.SendSuccess(c, http.StatusOK, createdUser, "User created successfully")
}

func getBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}

func GetUserHandler(c *gin.Context) {

	user, err := GetUserService()
	if err != nil {
		config.SendError(c, http.StatusInternalServerError, "Failed to retrieve user")
		return
	}

	if user == nil {
		config.SendError(c, http.StatusFound, "User not found")
		return
	}

	config.SendSuccess(c, http.StatusOK, user, "User retrieved successfully")
}
