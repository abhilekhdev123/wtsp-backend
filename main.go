package main

import (
	"log"
	"os"
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
	log.Println(`Server started on port`, os.Getenv("PORT"))
	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	r.Run(":" + os.Getenv("PORT")) // Start the server on the specified port
}
