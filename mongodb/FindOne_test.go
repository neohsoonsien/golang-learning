package mongodb

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestFindOne(t *testing.T) {
	t.Log("TestFindOne begins")
	filter := bson.D{{"name", "James"}}
	student := FindOne(mongoServer.URI(), filter)
	t.Logf("The foun student is %v", student)
}
