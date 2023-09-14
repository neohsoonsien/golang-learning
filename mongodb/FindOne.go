package mongodb

import (
	"context"
	"fmt"
	"os"

	godotenv "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func FindOne() {

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

	// ******************************************************* //
	// Step 3: Query from the database
	// ******************************************************* //
	// query from the collection
	collection := client.Database("tracker").Collection("coindex1_nft_price")
	opts := options.FindOne().SetSort(bson.D{{"tran_price", 1}})
	nftPrice := &Coindex1NftPrice{}
	err = collection.FindOne(context.TODO(), bson.D{{"token_id", "388005071993967880"}}, opts).Decode(nftPrice)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return
		}
		logger.Error(err)
	}
	logger.Infof("The found NFT is %v", nftPrice)
}
