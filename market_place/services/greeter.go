package services

import (
	"fmt"

	proto "golang-learning/market_place/proto"
	"golang-learning/market_place/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type GreeterService struct {
	// Could include dependencies like repositories, configs, etc.
}

func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

// business logics
func (s *GreeterService) GenerateGreeting(name string) (string, error) {
	// validation and business rules, etc.
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}

	return "Hello " + name, nil
}

func (s *GreeterService) PageInfo(pageRequest *proto.SearchRequest, totalItem int32) (*proto.PageInfo, error) {
	if totalItem == 0 {
		return nil, fmt.Errorf("totalItem cannot be empty")
	}

	// extract PageSize and CurrentPageNumber
	pageSize := int32(10)
	currentPageNumber := int32(1)
	if pageRequest != nil {
		pageSize = pageRequest.PageSize
		currentPageNumber = pageRequest.CurrentPageNumber
	}

	// calculate TotalPageNumber
	totalPageNumber := int32(1)
	if int32(totalItem%pageSize) == 0 {
		totalPageNumber = int32(totalItem / pageSize)
	} else {
		totalPageNumber = int32(totalItem/pageSize) + 1
	}

	// calculate CurrentPageSize
	currentPageSize := pageSize
	if currentPageNumber >= totalPageNumber {
		currentPageSize = int32(totalItem % pageSize)
		currentPageNumber = totalPageNumber
	}

	return &proto.PageInfo{
		PageSize:          pageSize,
		CurrentPageNumber: currentPageNumber,
		TotalItem:         totalItem,
		CurrentPageSize:   currentPageSize,
		TotalPageNumber:   totalPageNumber,
	}, nil
}

func (s *GreeterService) ListInvoices(req *proto.ListInvoicesRequest) ([]*proto.Invoice, error) {
	mongodbClient, err := repositories.InitMongoDB("mongodb://username:password@127.0.0.1:27017/marketplace")
	fmt.Printf("MongoDB client: %v, error: %v", mongodbClient, err)

	invoiceRepository := repositories.NewInvoiceRepository(mongodbClient.Database(repositories.MARKETPLACE_DATABASE).Collection(repositories.INVOICES_COLLECTION))

	if req == nil {
		return nil, fmt.Errorf("ListInvoicesRequest cannot be empty")
	}

	// Creates a query filter to match documents in which the "number", "vendor", "dateTimeStart" match
	filter := bson.M{
		"number": req.GetNumber(),
		"vendor": req.GetVendor(),
		"dateTime": bson.M{
			"$gte": req.GetDateTimeStart(),
			"$lt":  req.GetDateTimeStart() + int32(24*3600), // one day time range
		},
	}

	invoices, err := invoiceRepository.ListInvoices(filter)
	if err != nil {
		return nil, err
	}

	pbInvoices := []*proto.Invoice{}
	for _, invoice := range *invoices {
		pbInvoices = append(pbInvoices, invoice.ToPb())
	}

	return pbInvoices, nil
}
