package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"policy-search-api/controllers"
	"policy-search-api/services"
	"time"

	_ "policy-search-api/docs" // Import docs for Swagger

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}

	// MongoDB connection setup
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	collection := client.Database("policydb").Collection("policies")

	// Check if the collection is empty and insert a sample document if so
	count, err := collection.EstimatedDocumentCount(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		fmt.Println("Collection is empty, inserting a sample policy...")
		_, err := collection.InsertOne(ctx, bson.M{
			"_id":            "policy123",
			"name":           "Health Insurance",
			"startdate":      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			"durationyears":  5,
			"company":        "HealthCorp",
			"initialdeposit": 1000.0,
			"type":           "Health",
			"usertypes":      []string{"Individual", "Family"},
			"termsperyear":   4,
			"termamount":     250.0,
			"interest":       2.5,
		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Collection is not empty. Skipping insertion.")
	}

	// Initialize services
	policyService := services.NewPolicyService(collection)

	// Register routes
	r := mux.NewRouter()
	controllers.RegisterPolicyRoutes(r, policyService)

	// Swagger documentation route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
