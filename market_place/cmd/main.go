// package main

// import (
//   "net/http"

//   "github.com/gin-gonic/gin"
// )

// func main() {
//   // Create a Gin router with default middleware (logger and recovery)
//   r := gin.Default()

//   // Define a simple GET endpoint
//   r.GET("/ping", func(c *gin.Context) {
//     // Return JSON response
//     c.JSON(http.StatusOK, gin.H{
//       "message": "pong",
//     })
//   })

//   // Start server on port 8080 (default)
//   // Server will listen on 0.0.0.0:8080 (localhost:8080 on browser)
//   r.Run()
// }

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "golang-learning/market_place/pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8081, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
