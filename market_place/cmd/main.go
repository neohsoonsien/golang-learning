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

// package main

// import (
// 	"context"
// 	"flag"
// 	"net/http"

// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// 	"google.golang.org/grpc/grpclog"

// 	gw "golang-learning/market_place/proto" // Update
// )

// var (
// 	// command-line options:
// 	// gRPC server endpoint
// 	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
// )

// func run() error {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	// Register gRPC server endpoint
// 	// Note: Make sure the gRPC server is running properly and accessible
// 	mux := runtime.NewServeMux()
// 	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
// 	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
// 	if err != nil {
// 		return err
// 	}

// 	// Start HTTP server (and proxy calls to gRPC server endpoint)
// 	return http.ListenAndServe(":8081", mux)
// }

// func main() {
// 	flag.Parse()

// 	if err := run(); err != nil {
// 		grpclog.Fatal(err)
// 	}
// }

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"

	"golang-learning/market_place/controllers"
	pb "golang-learning/market_place/proto"
	"golang-learning/market_place/services"
)

var (
	// command-line options:
	grpcPort = flag.Int("grpc_port", 8080, "The gRPC server port")
	httpPort = flag.Int("http_port", 8081, "The http server port")
)

func grpcServer() {
	flag.Parse()

	// listening to port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// services and controllers dependencies injection
	greeterService := services.NewGreeterService()
	greeterController := controllers.NewGreeterController(greeterService)

	// start the server
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, greeterController)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func httpGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// gRPC server endpoint
	grpcServerEndpoint := fmt.Sprintf("localhost:%v", *grpcPort)

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf(":%v", *httpPort), mux)
}

func main() {
	flag.Parse()

	// start the gRPC server
	go grpcServer()

	// concurrent reverse-proxy for the http-gateway
	if err := httpGateway(); err != nil {
		grpclog.Fatal(err)
	}
}
