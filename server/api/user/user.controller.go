package user

import (
	"net/http"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	createdUser, err := CreateUserService(user)

	if err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		config.SendError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	// ... rest of the code
	config.SendSuccess(c, http.StatusOK, createdUser, "User created successfully")
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
