package mongodb

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	godotenv "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type Coindex1NftPrice struct {
	Id      primitive.ObjectID `bson:"_id"`
	TokenId string             `bson:"token_id"`
	Address string             `bson:"address"`
	Price   string             `bson:"tran_price"`
}

func Find_CompileScript() {

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
	// Step 3: Read in from the csv file
	// ******************************************************* //
	// read and parse the contents from a csv file
	file, err := os.Open("tokenId_wallet.csv")
	if err != nil {
		logger.Error(err)
	}
	csv := csv.NewReader(file)
	tokenIdList := make([]string, 0)
	tokenIdAddressMap := make(map[string]Coindex1NftPrice)
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error(err)
		}
		// input the data as map
		for index := range record {
			if index == 0 {
				tokenIdList = append(tokenIdList, record[0])
				logger.Infof("Token Id is %v", record[0])
			}
			tokenIdAddressMap[record[0]] = Coindex1NftPrice{
				TokenId: record[0],
				Address: record[1],
				Price:   "0",
			}
		}
	}
	logger.Infof("The number fo tokenIds is %v", len(tokenIdList))

	// ******************************************************* //
	// Step 4: Query from the database
	// ******************************************************* //
	// query from the collection
	collection := client.Database("tracker").Collection("coindex1_nft_price")
	opts := options.Find().SetSort(bson.D{{"tran_price", 1}})
	cursor, err := collection.Find(context.TODO(), bson.D{{"token_id", bson.D{{"$in", tokenIdList}}}}, opts)
	if err != nil {
		logger.Error(err)
	}
	// decode the cursor
	nftPriceList := make([]*Coindex1NftPrice, 0)
	if err = cursor.All(context.TODO(), &nftPriceList); err != nil {
		logger.Error(err)
	}
	logger.Infof("The number of nftPrice %v", len(nftPriceList))

	// ******************************************************* //
	// Step 5: Update the map using the results from the mongodb query
	// ******************************************************* //
	for _, nftPrice := range nftPriceList {
		if tokenIdAddress, exist := tokenIdAddressMap[nftPrice.TokenId]; exist {
			tokenIdAddressMap[nftPrice.TokenId] = Coindex1NftPrice{
				TokenId: tokenIdAddress.TokenId,
				Address: tokenIdAddress.Address,
				Price:   nftPrice.Price,
			}
		}
	}

	// ******************************************************* //
	// Step 6: Concatenate the content and write compile into javascript
	// ******************************************************* //
	// Concatenate the content into a string
	content := ""
	for _, value := range tokenIdAddressMap {
		if value.Price == "0" {
			continue
		}
		content += ("db.getCollection(\"tracker\").updateOne({\"nftnodes.tokenid\":\"" + value.TokenId + "\",\"address\":\"" + value.Address + "\"},{\"$set\":{\"nftnodes.$[nftnodes].cost\":\"" + value.Price + "\"}},{arrayFilters: [{\"nftnodes.tokenid\":\"" + value.TokenId + "\"}] })\n")
	}
	// write the content into a javascript
	data := []byte(content)
	_ = os.WriteFile("script.js", data, 0644)
}
