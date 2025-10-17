package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Question model for response
type Question struct {
	Question   string `json:"question"`
	QuestionId string `json:"questionId"`
	Answers    []struct {
		AnswerId string `json:"answerId"`
		Answer   string `json:"answer"`
	} `json:"answers"`
}

type QuestionProvider interface {
	GetQuestionById(ctx context.Context, questionId string) (*Question, error)
}

// Real Mongo provider
type MongoQuestionProvider struct {
	Coll *mongo.Collection
}

func (m *MongoQuestionProvider) GetQuestionById(ctx context.Context, questionId string) (*Question, error) {
	var raw struct {
		QuestionId string `bson:"questionId"`
		Question   string `bson:"question"`
		Answers    []struct {
			Text      string `bson:"text"`
			IsCorrect bool   `bson:"isCorrect"`
		} `bson:"answers"`
	}
	if err := m.Coll.FindOne(ctx, bson.M{"questionId": questionId}).Decode(&raw); err != nil {
		return nil, err
	}
	answers := make([]struct {
		AnswerId string `json:"answerId"`
		Answer   string `json:"answer"`
	}, len(raw.Answers))
	for i, a := range raw.Answers {
		answers[i].AnswerId = fmt.Sprintf("%02d", i+1)
		answers[i].Answer = a.Text
	}
	return &Question{
		Question:   raw.Question,
		QuestionId: raw.QuestionId,
		Answers:    answers,
	}, nil
}

// Handler using provider
func GetQuestionHandlerWithProvider(provider QuestionProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		questionId := c.Query("questionId")
		if questionId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing questionId parameter"})
			return
		}
		ctx := c.Request.Context()
		q, err := provider.GetQuestionById(ctx, questionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Question not found"})
			return
		}
		c.JSON(http.StatusOK, q)
	}
}

// Default handler for production
func GetQuestionHandler(c *gin.Context) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017/sre_quiz"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database connection error"})
		return
	}
	questionsColl := client.Database("sre_quiz").Collection("questions")
	provider := &MongoQuestionProvider{Coll: questionsColl}
	GetQuestionHandlerWithProvider(provider)(c)
}
