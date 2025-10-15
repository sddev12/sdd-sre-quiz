package db

import (
	"context"
	"log"
	"net/url"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance    *mongo.Client
	clientInstanceErr error
	clientMutex       sync.Mutex
)

// GetMongoClient returns a singleton MongoDB client instance
func GetMongoClient() (*mongo.Client, error) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	// If already initialized and no error, return the client
	if clientInstance != nil && clientInstanceErr == nil {
		return clientInstance, nil
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017/sre_quiz" // Example fallback connection string
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		clientInstanceErr = err
		return nil, err
	}
	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		clientInstanceErr = err
		return nil, err
	}
	// Redact credentials in URI for logging
	if parsed, err := url.Parse(uri); err == nil {
		log.Printf("Connected to MongoDB at %s", parsed.Redacted())
	} else {
		log.Printf("Connected to MongoDB (URI redaction failed)")
	}
	clientInstance = client
	clientInstanceErr = nil
	return clientInstance, nil
}
