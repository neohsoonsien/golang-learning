package pagination

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("BEFORE the tests!")
	exitVal := m.Run()
	log.Print("AFTER the tests!")

	os.Exit(exitVal)
}

func TestGetPageRequest(t *testing.T) {
	pageNum := int32(20)
	pageSize := int32(10)

	pageReq := getPageRequest(pageNum, pageSize)

	log.Printf("The page number is %v", pageReq.GetPageNum())
	log.Printf("The page size is %v", pageReq.GetPageSize())
}
