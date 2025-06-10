package user

import (
	"net/http"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	req := c.MustGet("validatedBody").(*CreateUserRequest)

	createdUser, status, msg, err := CreateUserService(req)
	if err != nil || status != http.StatusOK {
		config.SendError(c, status, msg)
		return
	}

	config.SendSuccess(c, http.StatusOK, createdUser, msg)
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
