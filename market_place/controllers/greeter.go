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
