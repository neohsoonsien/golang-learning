package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type BulkWriteUpdateList struct {
	Filter bson.D
	Update bson.D
}

func BulkWrite(collection *mongo.Collection, bulkUpdateList []BulkWriteUpdateList) (int64, int64, int64, int64) {
	// initialize logger
	logger := zap.NewExample().Sugar()
	defer logger.Sync()

	// collect and compile the mongo WriteModel
	models := make([]mongo.WriteModel, 0)
	for _, bulkUpdate := range bulkUpdateList {
		models = append(models, mongo.NewUpdateOneModel().SetFilter(bulkUpdate.Filter).SetUpdate(bulkUpdate.Update).SetUpsert(false))
	}

	// perform the BulkWrite operations
	opts := options.BulkWrite().SetOrdered(false)
	res, err := collection.BulkWrite(context.TODO(), models, opts)
	if err != nil {
		logger.Fatal(err)
		return int64(0), int64(0), int64(0), int64(0)
	}

	return res.InsertedCount, res.DeletedCount, res.MatchedCount, res.ModifiedCount
}
