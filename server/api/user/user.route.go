package user

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := "users"
	rg.POST(userroute+"/sign-up", CreateUserHandler)
	rg.GET(userroute+"/", GetUserHandler)
}
