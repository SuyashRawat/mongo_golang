package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/SuyashRawat/mongo_golang/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialize the router
	r := httprouter.New()

	// Get MongoDB client and pass it to the controller
	client := getMongoClient()
	us := controllers.NewUserController(client)

	// Define routes
	r.GET("/user/:id", us.GetUser)
	r.POST("/user", us.CreateUser)
	r.DELETE("/user/:id", us.DeleteUser)

	// Start the HTTP server
	log.Println("Server running at http://localhost:9000")
	http.ListenAndServe("localhost:9000", r)
}

// getMongoClient initializes and returns a MongoDB client
func getMongoClient() *mongo.Client {
	// MongoDB URI
	uri := "mongodb://localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB successfully")
	return client
}
