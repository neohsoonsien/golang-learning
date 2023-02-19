package mongodb

import (
	"context"
	"fmt"
	"os"
	"encoding/json"

	godotenv "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func Find() {

	// ******************************************************* //
	// Step 1: obtain the mongoURI connection string from .env file
	// ******************************************************* //
	godotenv.Load()
	mongoURI := os.Getenv("MONGODB_URI")

	// initialize logger
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

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
	fmt.Println("Successfully connected and pinged to Find.")


	collection := client.Database("tracker").Collection("tracker_history")
	filter := bson.D{{"address", "0x5859AADB52d93D6AadA9D07301151443a866Cc4D"}}

	// to allow higher tolerance in memory usage
	enableDiskUse := true
	opt := &options.FindOptions{
		AllowDiskUse: &enableDiskUse,
	}

	cursor, err := collection.Find(context.TODO(), filter, opt)
	if err != nil {
		logger.Errorf("Failed to find address for nodes: %v, with error %v", filter, err)
	}

	// declare results, parse the cursor into &results
	var results []interface{}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// marshal indent to print out the result
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}