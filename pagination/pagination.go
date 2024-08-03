package pagination

import (
	pb "golang-learning/pagination/proto"
)

func getPageRequest(pageNum int32, pageSize int32) *pb.PageRequest {
	pageReq := &pb.PageRequest{
		PageNum:  pageNum,
		PageSize: pageSize,
	}

	return pageReq
}
