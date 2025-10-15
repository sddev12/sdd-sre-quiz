package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance    *mongo.Client
	clientInstanceErr error
	mongoOnce         sync.Once
)

// GetMongoClient returns a singleton MongoDB client instance
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
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
			return
		}
		// Ping the database to verify connection
		if err := client.Ping(ctx, nil); err != nil {
			clientInstanceErr = err
			return
		}
		log.Printf("Connected to MongoDB at %s", uri)
		clientInstance = client
	})
	return clientInstance, clientInstanceErr
}
