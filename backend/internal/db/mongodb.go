package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client holds the global MongoDB client instance
var Client *mongo.Client

// ConnectDB initializes the MongoDB connection and returns the client instance.
func ConnectDB() *mongo.Client {
	// Try to load the .env file. We don't fatal here because in production,
	// env variables might be set directly in the environment (e.g., Docker/Vercel).
	if err := godotenv.Load(); err != nil {
		log.Println("Info: No .env file found or error loading it. Relying on system environment variables.")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("FATAL: MONGO_URI environment variable is not set")
	}

	// Set a timeout of 10 seconds for the connection to establish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)

	// Attempt to connect
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("FATAL: Failed to connect to MongoDB: %v\n", err)
	}

	// Ping the database to verify the connection is active
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("FATAL: Failed to ping MongoDB. Connection might be unstable: %v\n", err)
	}

	fmt.Println("🚀 Successfully connected to MongoDB!")
	Client = client
	return client
}

// GetCollection is a helper function to retrieve strongly-typed collection references
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	// Defaulting the database name to "unionhub" based on project scope
	return client.Database("unionhub").Collection(collectionName)
}
