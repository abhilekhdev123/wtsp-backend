package user

import (
	"wtsp-backend/server/helpers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := "users"
	rg.POST(userroute+"/sign-up",
		helpers.ValidateBody(CreateUserRequest{}), // Middleware for validating the request body
		CreateUserHandler)
	rg.GET(userroute+"/", GetUserHandler) // Get all users , it should be protected
}
