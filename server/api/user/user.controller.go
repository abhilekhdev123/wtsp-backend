package user

import (
	"net/http"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	//req := c.MustGet("validatedBody").(*CreateUserRequest)

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, status, msg, err := CreateUserService(&req)
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
