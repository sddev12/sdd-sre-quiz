package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockQuestionProvider struct {
	questions map[string]*Question
}

func (m *mockQuestionProvider) GetQuestionById(ctx context.Context, questionId string) (*Question, error) {
	q, ok := m.questions[questionId]
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return q, nil
}

func setupQuestionRouterWithProvider(provider QuestionProvider) *gin.Engine {
	r := gin.New()
	r.GET("/question", GetQuestionHandlerWithProvider(provider))
	return r
}

func TestGetQuestionHandler_Unit(t *testing.T) {
	provider := &mockQuestionProvider{
		questions: map[string]*Question{
			"01": {
				Question:   "What is SRE?",
				QuestionId: "01",
				Answers: []struct {
					AnswerId string `json:"answerId"`
					Answer   string `json:"answer"`
				}{
					{AnswerId: "01", Answer: "Site Reliability Engineering"},
					{AnswerId: "02", Answer: "Software Resource Engineering"},
				},
			},
		},
	}
	r := setupQuestionRouterWithProvider(provider)

	t.Run("returns 200 and question for valid questionId", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/question?questionId=01", nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200 got %d, body: %s", w.Code, w.Body.String())
		}
	})

	t.Run("returns 400 for missing questionId", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/question", nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status 400 got %d", w.Code)
		}
	})

	t.Run("returns 500 for non-existent questionId", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/question?questionId=99", nil)
		r.ServeHTTP(w, req)
		if w.Code != http.StatusInternalServerError {
			t.Fatalf("expected status 500 got %d", w.Code)
		}
	})
}
