package auth

import (
	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes registers the authentication routes with the given router group.
func RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", Login)
}
