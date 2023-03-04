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

func Aggregate() {

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


	contractAddress := ""
	tokenId := []string{"434689068597515168", "434689045323322272", "434689598690431904"}
 	collection := client.Database("events").Collection("nfts")
	matchContract := bson.D{{"$match", bson.D{{"contract", "0x77f03b8e0a2F0d9D37a61503CBC0a3930663685C"}}}}
	matchTokenId := bson.D{{"$match", bson.D{{"tokenId", bson.D{{"$in", tokenId}}}}}}

	// to allow higher tolerance in memory usage
	enableDiskUse := true
	opt := &options.AggregateOptions{
		AllowDiskUse: &enableDiskUse,
	}

	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchContract, matchTokenId}, opt)
	if err != nil {
		logger.Errorf("Failed to find Contract Address: %v with error: %v", contractAddress, err)
	}

	// declare results, parse the cursor into &results
	var nfts []Nft
	if err = cursor.All(context.TODO(), &nfts); err != nil {
		panic(err)
	}

	if nfts == nil {
		logger.Infof("The contract could not be found in collection tracker")
	}

	// marshal indent to print out the nft
	for _, nft := range nfts {
		output, err := json.MarshalIndent(nft, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
		fmt.Printf("%v\n", nft.Contract)
	}
}