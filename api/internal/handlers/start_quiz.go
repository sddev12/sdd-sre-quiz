package handlers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sddev12/sdd-sre-quiz/api/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

type StartQuizRequest struct {
	Username string `json:"username" binding:"required,min=1"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func StartQuizHandler(c *gin.Context) {
	var req StartQuizRequest
	if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Username) == "" {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "invalid input"})
		return
	}

	client, err := db.GetMongoClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "database error"})
		return
	}
	usersColl := client.Database("sre_quiz").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if user already exists
	count, err := usersColl.CountDocuments(ctx, bson.M{"username": req.Username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, MessageResponse{Message: "user already exists"})
		return
	}

	// Insert new user
	_, err = usersColl.InsertOne(ctx, bson.M{
		"username":  req.Username,
		"createdAt": time.Now(),
		"answers":   []interface{}{},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "database error"})
		return
	}
	c.JSON(http.StatusCreated, MessageResponse{Message: "user created successfully"})
}
