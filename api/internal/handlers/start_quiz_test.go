package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// helper to create a test router with the StartQuizHandler
func setupRouter() *gin.Engine {
	r := gin.New()
	r.POST("/start-quiz", StartQuizHandler)
	return r
}

// helper to connect to test DB; skip tests if DB not available
func getTestDB(t *testing.T) *mongo.Client {
	t.Helper()
	uri := os.Getenv("MONGODB_TEST_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017/sre_quiz"
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Skipf("Skipping test: cannot connect to test MongoDB at %s: %v", uri, err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		t.Skipf("Skipping test: cannot ping test MongoDB at %s: %v", uri, err)
	}
	return client
}

func cleanupTestUser(t *testing.T, client *mongo.Client, username string) {
	t.Helper()
	_ = client.Database("sre_quiz").Collection("users").Drop(context.Background())
}

func TestStartQuizHandler(t *testing.T) {
	client := getTestDB(t)
	// ensure clean state
	cleanupTestUser(t, client, "")

	r := setupRouter()

	t.Run("returns 201 when new user created", func(t *testing.T) {
		payload := map[string]string{"username": "testuser1"}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/start-quiz", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusCreated {
			t.Fatalf("expected status %d got %d, body: %s", http.StatusCreated, w.Code, w.Body.String())
		}
	})

	t.Run("returns 400 when username missing", func(t *testing.T) {
		payload := map[string]string{}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/start-quiz", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("returns 400 when username already exists", func(t *testing.T) {
		// insert user first
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		usersColl := client.Database("sre_quiz").Collection("users")
		_, err := usersColl.InsertOne(ctx, bson.M{"username": "existinguser", "createdAt": time.Now(), "answers": []interface{}{}})
		if err != nil {
			t.Fatalf("failed to insert existing user: %v", err)
		}

		payload := map[string]string{"username": "existinguser"}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/start-quiz", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d got %d", http.StatusBadRequest, w.Code)
		}
	})

	// cleanup
	cleanupTestUser(t, client, "")
}
