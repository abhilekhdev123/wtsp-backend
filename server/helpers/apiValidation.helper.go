package helpers

import (
	"net/http"
	"reflect"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateBody1(schema interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyType := reflect.TypeOf(schema)
		var requestBody interface{}
		if bodyType.Kind() == reflect.Slice {
			requestBody = reflect.New(bodyType).Interface()
		} else {
			requestBody = reflect.New(bodyType).Interface()
		}
		if err := c.ShouldBindJSON(requestBody); err != nil {
			config.SendError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
			c.Abort()
			return
		}
		// Validate the struct
		if err := validate.Struct(requestBody); err != nil {
			config.SendError(c, http.StatusBadRequest, "Validation failed: "+err.Error())
			c.Abort()
			return
		}
		c.Set("validatedBody", requestBody)
		c.Next()
	}
}

func ValidateBody(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use reflection to make a new instance of the provided type
		body := reflect.New(reflect.TypeOf(obj)).Interface()

		if err := c.ShouldBindJSON(body); err != nil {
			config.SendError(c, http.StatusBadRequest, "Invalid request body1: "+err.Error())
			c.Abort()
			return
		}

		// Store in context
		c.Set("validatedBody", body)
		c.Next()
	}
}
