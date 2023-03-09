package mongodb

import (
	"context"
	"os"

	godotenv "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"fmt"
)

func UpdateMany() {

	// ******************************************************* //
	// Step 1: obtain the mongoURI connection string from .env file
	// ******************************************************* //
	godotenv.Load()
	mongoURI := os.Getenv("MONGODB_EVENTS_URI")

	// ******************************************************* //
	// Step 2: connect to the mongodb
	// ******************************************************* //
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	// ******************************************************* //
	// Step 3: UpdatetMany docs in the collection "nfts" 
	// ******************************************************* //
	coll := client.Database("events").Collection("nfts")
	filter := bson.D{{"contract", "0x77f03b8e0a2F0d9D37a61503CBC0a3930663685C"}}
	update := bson.D{{"$set", bson.D{{"contract", "0x06476427fa1DFe1a62447bdD1ba42fD5b9F65758"}}}}
	result, err := coll.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of matched documents: %d\n", result.MatchedCount)
	fmt.Printf("Number of documents modified: %d\n", result.ModifiedCount)
}