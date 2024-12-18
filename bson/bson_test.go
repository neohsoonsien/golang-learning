package bson

import (
	"log"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMain(m *testing.M) {
	log.Print("BEFORE the tests!")
	exitVal := m.Run()
	log.Print("AFTER the tests!")

	os.Exit(exitVal)
}

func TestMainRoutine(t *testing.T) {
	value, err := BsonToMap()
	if err != nil {
		t.Errorf("Error in getting the value from bson map.")
		return
	}
	assert.Equal(t, value, "value1")
}
