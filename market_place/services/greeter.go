package services

import "fmt"

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

func (s *GreeterService) PageInfo(pageRequest *pb.SearchRequest, totalItem int32) (*pb.PageInfo, error) {
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

	return &pb.PageInfo{
		PageSize:          pageSize,
		CurrentPageNumber: currentPageNumber,
		TotalItem:         totalItem,
		CurrentPageSize:   currentPageSize,
		TotalPageNumber:   totalPageNumber,
	}, nil
}
