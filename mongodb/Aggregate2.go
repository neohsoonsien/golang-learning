package mongodb
import (
	"context"
	"fmt"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type Proxy struct {
	ProxyId      primitive.ObjectID `bson:"_id"`
	BlockNumber  uint64             `json:"BlockNumber"`
	TxHash       string             `json:"TransactionHash" gencodec:"required"`
	Time         uint64             `json:"Time"`
	Address      string             `json:"Address"`
	ProxyAddress map[string]string  `json:"ProxyAddress"`
	CreatedTime  int64              `json:"CreatedTime"`
}

func Aggregate2() {

	// ******************************************************* //
	// Step 1: define the mongoURI connection string
	// ******************************************************* //
	mongoURI := ""

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
	matchAddress := bson.D{{"$match", bson.D{{"address", address}}}}
	newProxyAddress := make(map[string]string)
	newProxyAddress["0xA554690c3c7273Cd0F9c70Deb664109578AbaB02"] = "haha"
	updateProxyAddress := bson.D{{
		"$replaceWith", bson.D{{
			"$setField", bson.D{
				{ "field", "proxyaddress"},
				{ "input", "$$ROOT"},
				{ "value", bson.D{{"$mergeObjects", bson.A{"$proxyaddress", newProxyAddress} }} } } }} }}

	// to allow higher tolerance in memory usage
	enableDiskUse := true
	opt := &options.AggregateOptions{
		AllowDiskUse: &enableDiskUse,
	}

	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchAddress, updateProxyAddress}, opt)
	if err != nil {
		logger.Errorf("Failed to find Contract Address: %v with error: %v", address, err)
	}

	// declare results, parse the cursor into &results
	var proxies []Proxy
	if err = cursor.All(context.TODO(), &proxies); err != nil {
		panic(err)
	}

	if proxies == nil {
		logger.Infof("The contract could not be found in collection tracker")
		return
	}

	// marshal indent to print out the nft
	for _, proxy := range proxies {
		output, err := json.MarshalIndent(proxy, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
		fmt.Printf("%v\n", proxy.ProxyAddress)
	}
}