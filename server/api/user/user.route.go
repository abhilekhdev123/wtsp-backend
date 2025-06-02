package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/sign-up", CreateUserHandler)
}
