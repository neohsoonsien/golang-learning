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

func Match() {

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

	// the group key / value need to have {"key": "_id", "value": [{"key": "contract", "value": "$contract"}]}
 	collection := client.Database("events").Collection("nfts")
	matchStage := bson.D{{"$match", bson.D{{"contract", "0x06476427fa1DFe1a62447bdD1ba42fD5b9F65758"}}}}

	// to allow higher tolerance in memory usage
	enableDiskUse := true
	opt := &options.AggregateOptions{
		AllowDiskUse: &enableDiskUse,
	}

	// aggregation pipeline
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage}, opt)
	if err != nil {
		logger.Errorf("Failed to group contract address")
	}

	fmt.Println(cursor)

	// declare nfts, parse the cursor into &nfts
	var nfts []Nft
	if err = cursor.All(context.TODO(), &nfts); err != nil {
		panic(err)
	}

	// marshal indent to print out the nfts
	for _, nft := range nfts {
		cursor.Decode(&nft)
		output, err := json.MarshalIndent(nft, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	fmt.Printf("The number of documents is %v \n", len(nfts))
}