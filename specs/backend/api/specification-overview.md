# API Specification

The REST API for the SRE Quiz app will be as follows:

**Programming lannguage**: Go

**Web Framework**: Gin

**Logging**: Structured logging using the slog package

**Database Connection**: MongoDB using the Go MongoDB Driver package

**Dev Tools**: godotenv

**Deployment**: Containerised (Docker container)

## Description

The Rest API will handle the backend functionality for the SRE Quiz app.
It will connect to a MongoDB for data persistenace and retrieval.

In it's initial state it will not implement authentication but will be protected with cors confguration.

It will have rate limiting which, if possible will be implemented using a rate limiting middleware.

It will implement structured logging via a custom logger

Swagger specs will be produced as docs in yaml format

There will be a docs directroy in the code project with comprehensive docs for the service

## Example .env file

```env
MONGO_URI=mongodb://localhost:27017/sre_quiz
GIN_MODE=release
PORT=8080
CORS_ORIGINS=http://localhost:3000
LOG_LEVEL=info
RATE_LIMIT=100
```

## Example Go Struct Definitions

```go
// User represents a quiz participant
type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name      string             `bson:"name" json:"name"`
    CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
    UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// Question represents a quiz question
type Question struct {
    ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    QuestionID string             `bson:"questionId" json:"questionId"`
    Text       string             `bson:"text" json:"text"`
    Answers    []Answer           `bson:"answers" json:"answers"`
}

type Answer struct {
    Text      string `bson:"text" json:"text"`
    IsCorrect bool   `bson:"isCorrect" json:"isCorrect"`
}

// Attempt represents a user's quiz attempt
type Attempt struct {
    ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID          primitive.ObjectID `bson:"userId" json:"userId"`
    UserName        string             `bson:"userName" json:"userName"`
    Answers         []UserAnswer       `bson:"answers" json:"answers"`
    Score           int                `bson:"score" json:"score"`
    TotalTimeSeconds int               `bson:"totalTimeSeconds" json:"totalTimeSeconds"`
    StartedAt       time.Time          `bson:"startedAt" json:"startedAt"`
    CompletedAt     time.Time          `bson:"completedAt" json:"completedAt"`
}

type UserAnswer struct {
    QuestionID string `bson:"questionId" json:"questionId"`
    AnswerText string `bson:"answerText" json:"answerText"`
    IsCorrect  bool   `bson:"isCorrect" json:"isCorrect"`
}

// LeaderboardEntry (optional, can be derived from attempts)
type LeaderboardEntry struct {
    UserID          primitive.ObjectID `bson:"userId" json:"userId"`
    UserName        string             `bson:"userName" json:"userName"`
    Score           int                `bson:"score" json:"score"`
    TotalTimeSeconds int               `bson:"totalTimeSeconds" json:"totalTimeSeconds"`
    CompletedAt     time.Time          `bson:"completedAt" json:"completedAt"`
}
```

## Security: Hiding Correct Answers from the Frontend

- The API will never send the isCorrect field or any indicator of the correct answer to the frontend in the /question response.
- The frontend will only receive the question text and possible answers (answerId and text).
- The isCorrect field is used server-side only for answer evaluation.
- When a user submits an answer, the API will check correctness and return only the result (e.g., correct/incorrect, score), never the answer key.
- This prevents users from discovering correct answers by inspecting network traffic or frontend code.
