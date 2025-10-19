package mongodb

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func Ping() {

	// ******************************************************* //
	// Step 1: obtain the mongoURI connection string from .env file
	//         initialize the logger
	// ******************************************************* //
	godotenv.Load()
	mongoURI := os.Getenv("MONGODB_URI")

	// initialize logger
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	// ******************************************************* //
	// Step 2: create a new client and connect to the server
	// ******************************************************* //
	opts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// ******************************************************* //
	// Step 3: send a ping to confirm a successful connection
	// ******************************************************* //
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	logger.Infof("Pinged your deployment. You successfully connected to MongoDB!")
}
