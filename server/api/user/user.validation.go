package user

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CreateUserRequest defines the expected fields for user sign-up.
type CreateUserRequest struct {
	Email          string   `json:"email" binding:"required,email"`
	Name           string   `json:"name,omitempty"`
	Phone          string   `json:"phone,omitempty"`
	Role           []string `json:"role" binding:"required"`
	Password       string   `json:"password,omitempty"`
	ProfilePic     string   `json:"profilePic,omitempty"`
	DeviceID       string   `json:"deviceId,omitempty"`
	TargetPlatform string   `json:"targetPlatform,omitempty"`
	DeviceToken    string   `json:"deviceToken,omitempty"`
	IsGoogleLogin  *bool    `json:"isGoogleLogin,omitempty"`
}

// Custom validation for phone (optional, example)
var phoneRegex = regexp.MustCompile(`^\+?[0-9]{7,15}$`)

func UserValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		validate := validator.New()
		// Optional: custom validation for phone
		if req.Phone != "" && !phoneRegex.MatchString(req.Phone) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone format"})
			c.Abort()
			return
		}
		if err := validate.Struct(req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("validatedBody", req)
		c.Next()
	}
}
