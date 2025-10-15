package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/sddev12/sdd-sre-quiz/api/internal/db"
)

func main() {
	// Load environment variables from .env if present
	if err := godotenv.Load(); err != nil {
		log.Printf("Could not load .env file: %v", err)
	}

	// Connect to MongoDB
	if _, err := db.GetMongoClient(); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Example root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "SRE Quiz API is running",
		})
	})

	log.Printf("Starting server on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
