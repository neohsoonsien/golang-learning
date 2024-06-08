package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "golang-learning/grpc/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello, " + in.GetName().FirstName + in.GetName().LastName}, nil
}

// Classroom function
func (s *server) Classroom(ctx context.Context, in *pb.ClassroomRequest) (*pb.ClassroomResponse, error) {
	log.Printf("Student name: %v, student Id: %v", in.GetName(), in.GetId())
	// hard-code a classroom value
	classroom := "1T10"
	num := pb.StudentStatus_value["ACTIVE"]
	valueNum := pb.StudentStatus(num)
	valueString := pb.StudentStatus_ACTIVE.String()
	log.Printf("The key of the StudentStatus is %v \n", num)
	log.Printf("The StudentStatus is %v \n", valueNum)
	log.Printf("The StudentStatus is %v \n", valueString)

	return &pb.ClassroomResponse{Classroom: classroom}, nil
}

// Greeting function
func (s *server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {

	log.Println("User name: ", req.Username)
	log.Println("Country code: ", req.CountryCode)

	var greeting string

	switch req.CountryCode {
	case "us":
		greeting = "Hello " + req.Username
	case "mx":
		greeting = "Hola " + req.Username
	default:
		greeting = "Hola/Hello " + req.Username
	}

	return &pb.GreetResponse{
		Result: greeting,
	}, nil
}

func (s *server) Theater(ctx context.Context, req *pb.TheaterRequest) (*pb.TheaterResponse, error) {

	log.Println("Theater name is: ", req.Name)

	theater := &pb.TheaterResponse{
		Name:          req.Name,
		AvailableSeat: &pb.TheaterResponse_Count{Count: int32(10)},
	}

	// ... snip ...
	// Use a type switch to determine which oneof was set.
	switch response := theater.AvailableSeat.(type) {
	case *pb.TheaterResponse_Count: // Count is int32 type
	case *pb.TheaterResponse_ErrorLog: // ErrorLog is string type
		log.Printf("The response is %v", response)
	case nil:
	default:
		log.Printf("TheaterResponse has unexpected type %T", response)
	}

	return theater, nil
}

func (s *server) Book(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {

	log.Println("Book title is: ", req.Title)

	book := &pb.BookDetail{}
	if req != nil && req.Title != "" && req.Author != "" {
		book = &pb.BookDetail{
			Title:  req.Title,
			Author: req.Author,
			Year:   "2007",
			Price:  59.90,
		}
	}

	return &pb.BookResponse{
		Response: &pb.BookResponse_BookDetail{
			BookDetail: book,
		},
	}, nil
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
