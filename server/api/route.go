package api

import (
	"log"
	"net/http"
	"wtsp-backend/server/api/user"

	"github.com/gin-gonic/gin"
)

// all routes for the API
func Routes(r *gin.Engine) {

	// API versioning
	api := r.Group("/api/v1")

	// User routes
	userRoutes := api.Group("/user")
	user.RegisterUserRoutes(userRoutes)

	// Root route
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the WTSP Backend API",
		})
	})

	log.Println("API routes registered")
}
