package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getFromDb(ctx context.Context) (map[string]interface{}, error) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+MONGO_DB_USER+":"+MONGO_DB_PASS+"@localhost:3001"))

	if err != nil {
		return nil, err
	}

	collection := client.Database(MONGO_DB_NAME).Collection("products")

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var records []bson.M

	for cur.Next(ctx) {

		var record bson.M

		if err = cur.Decode(&record); err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	res := map[string]interface{}{}

	res = map[string]interface{}{
		"data": records,
	}

	return res, nil
}
