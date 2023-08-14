package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func FindOne_UpdateOne() {

	// ******************************************************* //
	// Step 1: define the mongoURI connection string
	// ******************************************************* //
	mongoURI := "mongodb://tracker:tracker!@localhost:28015/events?directConnection=true"

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

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	address := "0x0fc7343b1121381485f6283B234586B901a26451"
	collection := client.Database("events").Collection("proxies")
	findOneOpts := options.FindOne().SetSort(bson.D{{"_id", -1}})
	proxies := &Proxy{}
	err = collection.FindOne(
		context.TODO(),
		bson.D{{"address", address}},
		findOneOpts,
	).Decode(proxies)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			newProxyAddress := make(map[string]string)
			newProxyAddress["0xA554690c3c7273Cd0F9c70Deb664109578AbaB02"] = "0x0fc7343b1121381485f6283B234586B901a26451"
			res, err := collection.InsertOne(
				context.TODO(),
				bson.D{{"_id", primitive.NewObjectID()},
					{"blocknumber", uint64(234567)},
					{"transactionhash", "0xc6547h5"},
					{"time", uint64(17738743)},
					{"address", address},
					{"proxyaddress", newProxyAddress},
					{"createdtime", uint64(1123454345)}})
			if err != nil {
				logger.Error(err)
			}
			logger.Infof("inserted document with ID %v\n", res.InsertedID)
		}
		logger.Error(err)
	} else {
		updateOpts := options.Update().SetUpsert(false)
		proxyAddress := proxies.ProxyAddress
		if _, exist := proxyAddress["0xA554690c3c7273Cd0F9c70Deb664109578AbaB02"]; !exist {
			proxyAddress["0xA554690c3c7273Cd0F9c70Deb664109578AbaB02"] = "0x0fc7343b1121381485f6283B234586B901a26451"
		} else {
			logger.Infof("The proxy address already exists.")
			return
		}
		update := bson.D{{"$set", bson.D{{"proxyaddress", proxyAddress}}}}
		filter := bson.D{{"address", address}}
		result, err := collection.UpdateOne(context.TODO(), filter, update, updateOpts)
		if err != nil {
			logger.Error(err)
		}

		if result.MatchedCount != 0 {
			logger.Infof("updated an existing document")
			return
		}
		if result.UpsertedCount != 0 {
			logger.Infof("inserted a new document with ID %v\n", result.UpsertedID)
		}

		logger.Infof("found document %v", proxies)
	}
}
