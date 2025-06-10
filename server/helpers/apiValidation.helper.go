package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"

	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
)

// var validate = validator.New()

// func ValidateBody(obj interface{}) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Use reflection to make a new instance of the provided type
// 		body := reflect.New(reflect.TypeOf(obj)).Interface()

// 		if err := c.ShouldBindJSON(body); err != nil {
// 			config.SendError(c, http.StatusBadRequest, "Invalid request body1: "+err.Error())
// 			c.Abort()
// 			return
// 		}

// 		// Store in context
// 		c.Set("validatedBody", body)
// 		c.Next()
// 	}
// }

func ValidateBody(obj interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the body into a buffer so it can be read twice
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}

		// Restore the original body for future reads
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Use reflection to get a pointer to the right type
		newObj := reflect.New(reflect.TypeOf(obj)).Interface()

		// Create a strict decoder
		decoder := json.NewDecoder(bytes.NewBuffer(bodyBytes))
		decoder.DisallowUnknownFields() // <-- This makes it strict!

		if err := decoder.Decode(newObj); err != nil {
			config.SendError(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
			c.Abort()
			return
		}

		// Store in context
		c.Set("validatedBody", newObj)
		c.Next()
	}
}
