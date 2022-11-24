package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

type Event struct {
	EventName   string
	EventParams map[string]string
	Contract    string
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:1qazZAQ!@localhost:27017/events"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("events").Collection("events")

	// find all documents in which the "sender" field is "0xwalletaddress"
	// specify the Sort option to sort the returned documents by age in ascending order
	// opts := options.Find().SetSort(bson.D{{"createdtime", 1}})
	cursor, err := collection.Find(context.TODO(), bson.D{{"methodname", "transfer"}})
	if err != nil {
		log.Fatal(err)
	}

	// doc, err := bson.Marshal(bson.D{{"methodname", "transfer"}, {"createdtime", 1}})

	// get a list of all returned documents and print them out
	// see the mongo.Cursor documentation for more examples of using cursors
	var results []TransactionEvent
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	var numSenderDiff uint = 0
	var numReceiverDiff uint = 0
	var numValueDiff uint = 0
	var eventSizeNotFulfill uint = 0
	for _, event := range results {
		var sender string = event.Sender
		var receiver string = event.MethodParams["recipient"]
		var amount string = event.Value
		var from string = event.Events[0].EventParams["from"]
		var to string = event.Events[0].EventParams["to"]
		var eventValue string = event.Events[0].EventParams["value"]

		if sender != from {
			numSenderDiff++
		} else if receiver != to {
			numReceiverDiff++
		} else if amount != eventValue {
			numValueDiff++
		}

		if len(event.Events) != 1 {
			eventSizeNotFulfill++
		}
	}
	fmt.Println("There are a total number of ", len(results), "LCUSD transactions.")
	fmt.Println("The sender and event from are different: ", numSenderDiff)
	fmt.Println("The receiver and event to are different: ", numReceiverDiff)
	fmt.Println("The amount and event value do not tally: ", numValueDiff)
	fmt.Println("The event.EventParams not equal to 1: ", eventSizeNotFulfill)
}
