package main
import (
	"context"
	"fmt"
	"os"

	godotenv "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Event struct {
	EventName   string
	EventParams map[string]string
	Contract    string
}

type TransactionEvent struct {
	EventId         primitive.ObjectID `bson:"_id"`
	BlockNumber     uint64             `json:"blockNumber"`
	TxHash          string             `json:"transactionHash" gencodec:"required"`
	Time            uint64             `json:"time"`
	Sender          string             `json:"sender"`
	Receiver        string             `json:"receiver"`
	Value           string             `json:"value"`
	MethodName      string             `json:"methodname"`
	MethodParams    map[string]string  `json:"methodparams"`
	Events          []Event            `json:"events"`
	Token           string             `json:"token"`
	TokenId         string             `json:"tokenid"`
	Wallets         []string           `json:"wallets"`
	CreatedTime     int64              `json:"CreatedTime"`
	TransactionType string             `json:"transactiontype"`
}

func main() {

	// ******************************************************* //
	// Step 1: obtain the mongoURI connection string from .env file
	// ******************************************************* //
	godotenv.Load()
	mongoURI := os.Getenv("MONGODB_URI")

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

	
	// ******************************************************* //
	// Step 3: finding methodname: mintWithURI from the mongodb
	// ******************************************************* //
	coll := client.Database("events").Collection("events")
	filter := bson.D{{"methodname", "mintWithURI"}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}

	var results []TransactionEvent
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	fmt.Println(len(results))

	// ******************************************************* //
	// Step 4: process the results into docs, 
	//		   which will be inserted into new collection "nfts"
	// ******************************************************* //
	docs := []interface{}{}
	for _, result := range results {
		docs = append(docs, bson.D{
							{"contract", result.Receiver}, 
							{"baseURI", "https://sevenseas.mypinata.cloud/ipfs/"},
							{"tokenURI", result.MethodParams["tokenURI"]},
							{"tokenId", result.TokenId},
							{"wallets", result.Wallets},
						})
	}

	fmt.Println(len(docs))

	// ******************************************************* //
	// Step 5: InsertMany docs into the new collection "nfts" 
	// ******************************************************* //
	coll = client.Database("events").Collection("nfts")
	result, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of documents inserted: %d\n", len(result.InsertedIDs))
}