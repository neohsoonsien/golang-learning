package controllers

import (
	"context"
	"log"

	pb "golang-learning/market_place/pb"
	"golang-learning/market_place/services"
)

type GreeterController struct {
	pb.UnimplementedGreeterServer
	greeterService *services.GreeterService
}

func NewGreeterController(greeterService *services.GreeterService) *GreeterController {
	return &GreeterController{
		greeterService: greeterService,
	}
}

// gRPC handlers
func (c *GreeterController) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())

	// Delegate business logic to service layer
	message, err := c.greeterService.GenerateGreeting(in.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.HelloResponse{Message: message}, nil
}

func (c *GreeterController) PageInfo(ctx context.Context, req *pb.PageInfoRequest) (*pb.PageInfoResponse, error) {
	log.Printf("Received: %v", req.GetSearchRequest())

	// Delegate business logic to service layer
	pageInfo, err := c.greeterService.PageInfo(req.GetSearchRequest(), req.GetTotalItem())
	if err != nil {
		return nil, err
	}

	return &pb.PageInfoResponse{
		PageInfo: pageInfo,
	}, nil
}
