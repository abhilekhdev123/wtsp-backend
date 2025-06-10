package servey

import (
	"wtsp-backend/server/helpers"

	"github.com/gin-gonic/gin"
)

func ServeyRoute(rg *gin.RouterGroup) {
	serveyroute := "survey"

	rg.POST(serveyroute,
		helpers.ValidateBody(CreateServeyRequest{}),
		CreateServeyHandler)

	rg.GET(serveyroute, GetServeyList)
}
