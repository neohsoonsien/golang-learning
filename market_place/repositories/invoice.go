package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

	"golang-learning/market_place/models"
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
	// defer func() error {
	// 	if err = client.Disconnect(context.Background()); err != nil {
	// 		return fmt.Errorf("failed to disconnect from MongoDB client, MongoDB URI: %v, error: %v", mongodbURI, err)
	// 	}
	// 	return nil
	// }()

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database(DATABASE).RunCommand(context.Background(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to ping connection to MongoDB client, MongoDB URI: %v, error: %v", mongodbURI, err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}

func ListInvoices(client *mongo.Client, filter bson.M) (*[]models.Invoice, error) {
	coll := client.Database(DATABASE).Collection(COLLECTION)

	// Retrieves documents that match the query filter
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to Find from %v, error: %v", COLLECTION, err)
	}

	// Unpacks the cursor into a slice
	var invoices []models.Invoice
	if err = cursor.All(context.Background(), &invoices); err != nil {
		return nil, fmt.Errorf("failed to unpack cursor into 'invoices' slices")
	}

	// Prints the invoices of the find operation as structs
	for _, invoice := range invoices {
		output, err := json.MarshalIndent(invoice, "", "    ")
		if err != nil {
			return nil, fmt.Errorf("failed to marshal the 'invoices'")
		}
		fmt.Printf("%s\n", output)
	}

	return &invoices, nil
}
