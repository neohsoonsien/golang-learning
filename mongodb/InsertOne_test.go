package mongodb

import (
	"log"
	"os"
	"testing"

	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gotest.tools/v3/assert"
)

var (
	err         error
	mongoServer *memongo.Server
)

func TestMain(m *testing.M) {
	log.Print("BEFORE the tests!")
	mongoServer, err = memongo.Start("6.0.4")
	if err != nil {
		log.Fatal(err)
	}
	defer mongoServer.Stop()

	exitVal := m.Run()
	log.Print("AFTER the tests!")

	os.Exit(exitVal)
}

func TestInsertOne(t *testing.T) {
	t.Log("TestInsertOne begins")

	// insert the data
	insert := &Student{
		Id:   primitive.NewObjectID(),
		Name: "Jason",
		Age:  23,
	}
	res := InsertOne(mongoServer.URI(), insert)
	log.Printf("The inserted student is %v", res)

	assert.DeepEqual(t, insert, res)
}
