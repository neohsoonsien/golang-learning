package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestNewInvoiceRepository(t *testing.T) {
	mockColl := &MockCollection{}

	repo := NewInvoiceRepository(mockColl)

	assert.NotNil(t, repo)
	assert.Equal(t, mockColl, repo.collection)
}

func TestListInvoices(t *testing.T) {
	mockColl := &MockCollection{}

	// Prepare test data
	testData := []interface{}{
		bson.M{
			"principalId": "user_2",
			"number":      "20250620_0002",
			"vendor":      "Vendor B",
			"dateTime":    1753637000,
			"details": bson.A{
				bson.M{
					"name":       "Item 3",
					"category":   "Veges",
					"unitPrice":  9,
					"weightUnit": "kg",
					"count":      3,
					"countUnit":  "kg",
				},
				bson.M{
					"name":       "Item 4",
					"category":   "Veges",
					"unitPrice":  7,
					"weightUnit": "",
					"count":      6,
					"countUnit":  "",
				},
			},
		},
	}

	// Create mock cursor with test data
	cursor := createMockCursor(t, testData)

	// Setup mock expectation
	mockColl.On("Find",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(cursor, nil)

	filter := bson.M{
		"number": "20250620_0002",
		"vendor": "Vendor B",
		"dateTime": bson.M{
			"$gte": 1753635000,
			"$lt":  1753635000 + int32(24*3600), // one day time range
		},
	}

	repo := NewInvoiceRepository(mockColl)
	repo.ListInvoices(filter)

	// Verify expectations were met
	mockColl.AssertExpectations(t)
}
