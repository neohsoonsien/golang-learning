package mongodb

import (
	"log"
	"os"
	"testing"

	"github.com/tryvium-travels/memongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	insert := bson.D{{"_id", primitive.NewObjectID()}, {"name", "Jason"}, {"age", 23}}
	InsertOne(mongoServer.URI(), insert)
}
