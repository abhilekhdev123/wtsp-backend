package main

import (
	"log"

	"wtsp-backend/server/api"
	"wtsp-backend/server/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	_ = godotenv.Load("server/config/.env")
	// Connect to DB
	config.ConnectDB()

	// Start Gin server
	r := gin.Default()

	// main API routes call
	api.Routes(r)

	// Start server
	log.Println("Server started on port 8080")
	r.Run(":8080")
}
