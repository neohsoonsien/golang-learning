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

	if setUpdate, setUpdateExist := object["$set"]; setUpdateExist {
		if fieldUpdate, fieldUpdateExist := setUpdate.(primitive.M)["field1"]; fieldUpdateExist {
			log.Printf("key: field1, value: %v", fieldUpdate)
			return fieldUpdate.(string), nil
		}
	}

	return "", errors.New("Failed to get bson map value.")
}
