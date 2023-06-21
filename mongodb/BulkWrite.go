package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.uber.org/zap"
)

func BulkWrite() {

	// ******************************************************* //
	// Step 1: define the mongoURI connection string
	// ******************************************************* //
	mongoURI := "mongodb://<username>:<password>@localhost:27017/events?directConnection=true"

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

	address := "0x0fc7343b1121381485f6283B234586B901a26451"
	collection := client.Database("events").Collection("proxies")
   	findOneOpts := options.FindOne().SetSort(bson.D{{"_id", -1}})
   	var proxies Proxy
   	err = collection.FindOne(
		context.TODO(),
	   	bson.D{{"address", address}},
	   	findOneOpts,
   	).Decode(&proxies)

	fmt.Println(proxies)

	// Update the "email" field for two users.
	// For each update, specify the Upsert option to insert a new document if a
	// document matching the filter isn't found.
	// Set the Ordered option to false to allow both operations to happen even
	// if one of them errors.
	firstUpdate := bson.D{
		{"$set", bson.D{
			{"txhash", "new hash"},
		}},
	}
	models := []mongo.WriteModel{
		mongo.NewUpdateOneModel().SetFilter(bson.D{{"address", "0x0fc7343b1121381485f6283B234586B901a26451"}}).
			SetUpdate(firstUpdate).SetUpsert(true),
	}
	opts := options.BulkWrite().SetOrdered(false)
	res, err := collection.BulkWrite(context.TODO(), models, opts)
	if err != nil {
		logger.Fatal(err)
	}

	fmt.Printf(
		"inserted %v and deleted %v documents\n",
		res.InsertedCount,
		res.DeletedCount)
}