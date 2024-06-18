package mongodb

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/v3/assert"
)

func TestFindOne(t *testing.T) {
	t.Log("TestFindOne begins")

	// insert the data
	insert := &Student{
		Id:   primitive.NewObjectID(),
		Name: "James",
		Age:  25,
	}
	res := InsertOne(mongoServer.URI(), insert)

	// retrieve the data
	filter := bson.D{{"name", "James"}}
	student := FindOne(mongoServer.URI(), filter)
	t.Logf("The found student is %v", student)

	assert.DeepEqual(t, student, res)
}
