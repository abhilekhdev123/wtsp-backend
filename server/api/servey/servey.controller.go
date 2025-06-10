package servey

import (
	"net/http"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
)

func CreateServeyHandler(c *gin.Context) {
	req := c.MustGet("validatedBody").(*CreateServeyRequest)

	createdUser, status, msg, err := CreateServeyService(req)
	if err != nil || status != http.StatusOK {
		config.SendError(c, status, msg)
		return
	}

	config.SendSuccess(c, http.StatusOK, createdUser, msg)

}

func GetServeyList(c *gin.Context) {

	serveylist, err := GetServeyListFromService()
	if err != nil {
		config.SendError(c, http.StatusInternalServerError, "Error getting servey list")
		return
	}

	if serveylist == nil {
		config.SendError(c, http.StatusNotFound, "No servey list found")
	}

	config.SendSuccess(c, http.StatusOK, serveylist, "Servey list retrieved successfully")

}
