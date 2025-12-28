package mongodb

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gotest.tools/v3/assert"
)

func TestBulkWrite(t *testing.T) {

	t.Log("TestBulkWrite begins")

	// ******************************************************* //
	// Step 1: define the mongoURI connection string
	// ******************************************************* //
	mongoURI := "mongodb://gemini:password@localhost:27017/gemini?directConnection=true"

	// ******************************************************* //
	// Step 2: connect to the mongodb
	// ******************************************************* //
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			t.Error(err)
		}
	}()
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Error(err)
	}
	t.Log("Successfully connected and pinged.")

	// ******************************************************* //
	// Step 3: insert many to mongodb collection
	// ******************************************************* //
	collection := client.Database("gemini").Collection("bulkWrites")
	res, err := collection.InsertMany(context.TODO(), bson.A{
		bson.D{{"_id", primitive.NewObjectID()}, {"user", "user1"}, {"amount", 20}},
		bson.D{{"_id", primitive.NewObjectID()}, {"user", "user2"}, {"amount", 30}},
		bson.D{{"_id", primitive.NewObjectID()}, {"user", "user3"}, {"amount", 40}},
		bson.D{{"_id", primitive.NewObjectID()}, {"user", "user4"}, {"amount", 50}},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("The inserted documents' ID are: %v", res.InsertedIDs)

	// ******************************************************* //
	// Step 4: test BulkWrite into the collection
	// ******************************************************* //
	bulkUpdateList := []BulkWriteUpdateList{
		{Filter: bson.D{{"user", "user1"}}, Update: bson.D{{"$set", bson.D{{"amount", 50}}}}},
		{Filter: bson.D{{"user", "user2"}}, Update: bson.D{{"$set", bson.D{{"amount", 40}}}}},
		{Filter: bson.D{{"user", "user3"}}, Update: bson.D{{"$set", bson.D{{"amount", 30}}}}},
		{Filter: bson.D{{"user", "user4"}}, Update: bson.D{{"$set", bson.D{{"amount", 20}}}}},
	}
	insertedCount, deletedCount, matchedCount, modifiedCount := BulkWrite(collection, bulkUpdateList)

	// ******************************************************* //
	// Step 5: remove all the documents from the collection
	// ******************************************************* //
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		t.Error(err)
	}
	if deleteResult.DeletedCount != int64(4) {
		t.Errorf("Expect to delete ['4'] documents, but deleted only ['%v']", deleteResult.DeletedCount)
	}

	// ******************************************************* //
	// Step 6: verify the results
	// ******************************************************* //
	assert.Equal(t, insertedCount, int64(0))
	assert.Equal(t, deletedCount, int64(0))
	assert.Equal(t, matchedCount, int64(4))
	assert.Equal(t, modifiedCount, int64(4))
}
