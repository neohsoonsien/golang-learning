package repositories

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const DATABASE = "marketplace"
const COLLECTION = "invoices"

func InitMongoDB(mongodbURI string) (*mongo.Client, error) {

	if mongodbURI == "" {
		log.Fatal("Please provide 'mongodbURI' variable. See\n\t https://docs.mongodb.com/drivers/go/current/usage-examples/")
	}

	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(mongodbURI).SetServerAPIOptions(serverAPI)

	// Creates a new client and connects to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB client, MongoDB URI: %v, error: %v", mongodbURI, err)
	}
	defer func() error {
		if err = client.Disconnect(context.Background()); err != nil {
			return fmt.Errorf("failed to disconnect from MongoDB client, MongoDB URI: %v, error: %v", mongodbURI, err)
		}
		return nil
	}()

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database(DATABASE).RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to ping connection to MongoDB client, MongoDB URI: %v, error: %v", mongodbURI, err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
