package config

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse defines the structure of error responses.
type ErrorResponse struct {
	Error   string `json:"error"`
	Details any    `json:"details,omitempty"`
}

// SendError sends a JSON error response with the given status code and message.
func SendError1(c *gin.Context, status int, message string, details ...any) {
	resp := ErrorResponse{Error: message}
	if len(details) > 0 {
		resp.Details = details[0]
	}
	c.AbortWithStatusJSON(status, resp)
}

// Example usage in a controller or middleware:
// config.SendError(c, http.StatusBadRequest, "Invalid input", err.Error())
