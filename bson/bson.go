package bson

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BsonToMap() (string, error) {
	set := bson.M{"field1": "value1"}
	unset := bson.M{"field2": "value2"}

	object := bson.M{}
	object["$set"] = set
	object["$unset"] = unset

	for mongoOperation, updates := range object {
		if mongoOperation == "$set" {
			for field, value := range updates.(primitive.M) {
				if field == "field1" {
					log.Printf("key: %v, value: %v", field, value)
					return value.(string), nil
				}
			}
		}
	}
	return "", errors.New("Failed to get bson map value.")
}
