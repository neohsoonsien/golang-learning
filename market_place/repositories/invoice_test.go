package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) Find(ctx context.Context, filter any, opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter any, opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*mongo.SingleResult)
}

func createMockCursor(t *testing.T, documents []interface{}) *mongo.Cursor {
	// Create cursor from BSON data
	cursor, err := mongo.NewCursorFromDocuments(documents, nil, nil)
	require.NoError(t, err)

	return cursor
}
}
