package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func ObjectID(mongoURI string, insert *Student) *Student {

	// ******************************************************* //
	// Step 1: initialize logger
	// ******************************************************* //
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
	collection := client.Database("students").Collection("students")
	res, err := collection.InsertOne(context.TODO(), insert)
	if err != nil {
		logger.Error(err)
	}

	logger.Infof("inserted document with ID %v\n", res.InsertedID)

	if res.InsertedID != primitive.NilObjectID {
		return insert
	}
	return nil
}
