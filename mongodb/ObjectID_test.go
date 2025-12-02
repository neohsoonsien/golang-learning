package mongodb

import (
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/v3/assert"
)

func TestTimestampObjectID(t *testing.T) {
	t.Log("TestTimestampObjectID begins")

	// insert the data
	insert := &Student{
		Id:   primitive.NewObjectID(),
		Name: "Jason",
		Age:  23,
	}
	res := ObjectID(mongoServer.URI(), insert)
	log.Printf("The inserted student is %v", res)

	// find the data
	find := FindOne(mongoServer.URI(), bson.D{{"_id", insert.Id}})
	currentTimestamp := find.Id.Timestamp()
	log.Printf("The current 'Timestamp' is %v", currentTimestamp)

	// extract timestamp from an existing ObjectID
	previousObjectID, _ := primitive.ObjectIDFromHex("66bc024fa781c1bf27ade3f7")
	previousTimestamp := previousObjectID.Timestamp()
	log.Printf("The previous 'Timestamp' is %v", previousTimestamp)

	// compare and verify the result
	log.Printf("The 'Timestamp' comparison is %v", previousTimestamp.Compare(currentTimestamp))
	assert.DeepEqual(t, previousTimestamp.Compare(currentTimestamp), -1)
}

func TestCompareObjectID(t *testing.T) {
	t.Log("TestCompareObjectID begins")

	// generate 2 different ObjectID
	firstObjectId := primitive.NewObjectID()
	secondObjectId := primitive.NewObjectID()

	assert.Assert(t, firstObjectId != secondObjectId)

	// generate same ObjectID
	secondObjectId = firstObjectId

	assert.Equal(t, firstObjectId, secondObjectId)
}

func TestNilObjectId(t *testing.T) {
	objectId := &primitive.NilObjectID

	assert.Equal(t, *objectId, primitive.NilObjectID)
}
