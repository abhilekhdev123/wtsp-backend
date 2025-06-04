package config

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(code, gin.H{
		"success": true,
		"message": message,
		"data":    data,
		"status":  code,
	})
}

func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"success": false,
		"message": message,
		"status":  code,
	})
}
