package redis

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"gotest.tools/v3/assert"
)

func TestGetStudent(t *testing.T) {
	t.Log("TestGetStudent begins")

	miniredis := miniredis.RunT(t)

	// set a key: value pair
	json, err := json.Marshal(Student{Name: "Adam", Id: "C4321"})
	if err != nil {
		log.Printf("Error in json encoding: %v", err)
	}
	err = miniredis.Set("C4321", string(json))
	if err != nil {
		log.Printf("Error in setting the key-value: %v", err)
	}

	// get the key: value
	result, err := miniredis.Get("C4321")
	if err != nil {
		log.Printf("Error in getting the key-value: %v", err)
	}
	log.Printf("The obtained result is %v", result)

	assert.DeepEqual(t, result, string(json))
}
