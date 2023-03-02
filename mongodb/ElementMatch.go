package mongodb

import (
	"context"
	"fmt"
	"os"
	"encoding/json"

	godotenv "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type TrackerNode struct {
	NodeId          primitive.ObjectID `bson:"_id"`
	Address         string             `json:"Address"`
	Generation      uint64             `json:"Generation"`
	BlockNumber     uint64             `json:"BlockNumber"`
	TxHash          string             `json:"TxHash"`
	Time            uint64             `json:"Time"`
	OperationType   string             `json:"OperationType"`
	LcusdNodes      []LcusdNode        `json:"LcusdNodes"`
	NftNodes        []NftNode          `json:"NftNodes"`
	BidNftNodes     []NftNode          `json:"BidNftNodes"`
	ConversionNodes []ConversionNode   `json:"ConversionNodes"`
	CreatedTime     int64              `json:"CreatedTime"`
}

type ParentNode struct {
	Generation uint64             `json:"Generation"`
	Id         primitive.ObjectID `json:"Id"`
}

type AgeAmount struct {
	Amount string `json:"Amount"`
	Age    uint64 `json:"Age"`
}

type LcusdNode struct {
	NodeId  primitive.ObjectID `bson:"_id"`
	Parents []ParentNode       `json:"Parents"`
	Amount  string             `json:"Amount"`
	Age     uint64             `json:"Age"`
}

type NftNode struct {
	NodeId   primitive.ObjectID `bson:"_id"`
	Parents  []ParentNode       `json:"Parents"`
	Exchange string             `json:"Exchange"`
	OrderId  string             `json:"OrderId"`
	Token    string             `json:"Token"`
	TokenId  string             `json:"TokenId"`
	Cost     string             `json:"Cost"`
	Amount   []AgeAmount        `json:"Amount"`
}

type ConversionNode struct {
	NodeId  primitive.ObjectID `bson:"_id"`
	Parents []ParentNode       `json:"Parents"`
	Time    uint64             `json:"Time"`
	Total   string             `json:"Total"`
	Amount  []AgeAmount        `json:"Amount"`
}

type Nft struct {
	NftId 			primitive.ObjectID `bson:"_id"`
	Contract     	string             `json:"contract"`
	BaseURI         string             `json:"baseURI"`
	TokenURI        string             `json:"tokenURI"`
	TokenId         string             `json:"tokenId"`
	Metadata        string             `json:"metadata"`
}

type Criteria struct {
	token			string
	tokenid			string
}

func ElementMatch() {

	// ******************************************************* //
	// Step 1: obtain the mongoURI connection string from .env file
	// ******************************************************* //
	godotenv.Load()
	mongoURI := os.Getenv("MONGODB_TRACKER_URI")

	// initialize logger
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	logger.Infof("mongoURI is : %s", mongoURI)

	// ******************************************************* //
	// Step 2: connect to the mongodb
	// ******************************************************* //
	// Create a new client and connect to the server
	clientTracker, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = clientTracker.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Ping the primary
	if err := clientTracker.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to tracker collection.")


	collection := clientTracker.Database("tracker").Collection("tracker_history")
	filter := bson.D{{"address", "0x5859AADB52d93D6AadA9D07301151443a866Cc4D"}}

	// sort in descending order to return only the first element in the FindOne query
	opt := &options.FindOneOptions{
		Sort:         bson.M{"time": -1},
	}

	// declare results, decode the cursor into &results
	var result *TrackerNode
	err = collection.FindOne(context.TODO(), filter, opt).Decode(&result)		
	if err != nil {
		panic(err)
	}

	if result == nil {
		logger.Infof("The wallet could not be found in collection tracker")
	}

	// marshal indent to print out the result
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
	fmt.Printf("%v\n", result.BlockNumber)

	for _, nftnode := range result.NftNodes {
		fmt.Printf("%v,\t%v\n", nftnode.Token, nftnode.TokenId)
	}


	collection = clientTracker.Database("tracker").Collection("tracker")
	enableDiskUse := true
	option := &options.FindOptions{
		AllowDiskUse: &enableDiskUse,
	}

	criteria := Criteria{
		token: "0xEA6EE9730609AC46D8d59F69ec576E27DDFE19C5",
		tokenid: "1461523938416383821135279884712579639185445486593",
	}

	filter = bson.D{{"nftnodes", bson.D{{"$elemMatch", criteria}}}}

	cursor, err := collection.Find(context.TODO(), filter, option)
	if err != nil {
		panic(err)
	}
	// end find
							
	var trackerNodes []TrackerNode
	if err = cursor.All(context.TODO(), &trackerNodes); err != nil {
		panic(err)
	}
							
	for _, tracker := range trackerNodes {
		cursor.Decode(&tracker)
		output, err := json.MarshalIndent(tracker, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
}