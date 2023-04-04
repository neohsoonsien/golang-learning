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

func Sort() {

	// ******************************************************* //
	// Step 1: obtain the mongoURI connection string from .env file
	// ******************************************************* //
	godotenv.Load()
	mongoURI := os.Getenv("MONGODB_EVENTS_URI")

	// initialize logger
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	logger.Infof("mongoURI is : %s", mongoURI)

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

	// sort in ascending order, and limit the first 5 documents only 
 	collection := client.Database("events").Collection("events")
	sortStage := bson.D{{"$sort", bson.D{{"time", 1}}}}
	limitStage := bson.D{{"$limit", 5}}

	// to allow higher tolerance in memory usage
	enableDiskUse := true
	opt := &options.AggregateOptions{
		AllowDiskUse: &enableDiskUse,
	}

	// aggregation pipeline
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{sortStage, limitStage}, opt)
	if err != nil {
		logger.Errorf("Failed to sort events")
	}

	fmt.Println(cursor)

	// declare events, parse the cursor into &events
	var events []TransactionEvent
	if err = cursor.All(context.TODO(), &events); err != nil {
		panic(err)
	}

	// marshal indent to print out the events
	for _, event := range events {
		cursor.Decode(&event)
		output, err := json.MarshalIndent(event, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	fmt.Printf("The number of events is %v \n", len(events))
}