package repositories

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestFindOperation(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("find documents", func(mt *mtest.T) {
		// Create a cursor response with sample data
		expected := bson.D{{"name", "John"}}

		// Mock responses for find operation:
		// 1. First batch with data
		firstBatch := mtest.CreateCursorResponse(1, "test.collection", mtest.FirstBatch, expected)
		// 2. Empty batch to signal end of cursor
		emptyBatch := mtest.CreateCursorResponse(0, "test.collection", mtest.NextBatch)

		// Add both responses - find() will consume them in sequence
		mt.AddMockResponses(firstBatch, emptyBatch)

		// Perform Find
		cursor, err := mt.Coll.Find(context.TODO(), bson.D{})
		if err != nil {
			mt.Fatalf("Find failed: %v", err)
		}
		defer cursor.Close(context.TODO())

		var results []bson.M
		if err := cursor.All(context.TODO(), &results); err != nil {
			mt.Fatalf("Cursor failed: %v", err)
		}

		// Assert
		if len(results) != 1 {
			mt.Fatalf("expected 1 result, got %d", len(results))
		}
		if results[0]["name"] != "John" {
			mt.Errorf("expected name 'John', got '%v'", results[0]["name"])
		}
	})
}
