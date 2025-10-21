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

type SubmitAnswerRequest struct {
	Username   string `json:"username" binding:"required,min=1"`
	QuestionId string `json:"questionId" binding:"required,min=1"`
	AnswerId   string `json:"answerId" binding:"required,min=1"`
}

func SubmitAnswerHandler(c *gin.Context) {
	var req SubmitAnswerRequest
	if err := c.ShouldBindJSON(&req); err != nil || strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.QuestionId) == "" || strings.TrimSpace(req.AnswerId) == "" {
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

	// Find user
	filter := bson.M{"username": req.Username}
	update := bson.M{
		"$push": bson.M{"answers": bson.M{"questionId": req.QuestionId, "answerId": req.AnswerId}},
	}
	result := usersColl.FindOne(ctx, filter)
	if result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			c.JSON(http.StatusNotFound, MessageResponse{Message: "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "database error"})
		return
	}

	// Update user record with answer
	_, err = usersColl.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, MessageResponse{Message: "database error"})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{Message: "answer recorded"})
}
