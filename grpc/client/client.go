package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "golang-learning/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// test SayHello function
	name := &pb.Name{FirstName: "John", LastName: "Wick"}
	sayHello, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", sayHello.GetMessage())

	// test Classroom function
	classroom, err := c.Classroom(ctx, &pb.ClassroomRequest{Name: "James", Id: "C2341"})
	if err != nil {
		log.Fatalf("could not find classroom: %v", err)
	}
	log.Printf("Classroom: %s", classroom.GetClassroom())

	// test Greeting function
	getGreeting("Jack", "us", c)
	getGreeting("Jose", "mx", c)

	// test Theater function
	getTheater("Golden Screen Cinema", c)

	// test Book function
	getBook("Harry Potter", c)
}

func getGreeting(name string, countryCode string, c pb.GreeterClient) {

	log.Println("Creating greeting")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		CountryCode: countryCode,
		Username:    name,
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)
}

func getTheater(name string, c pb.GreeterClient) {

	log.Println("Getting Theater")

	res, err := c.Theater(context.Background(), &pb.TheaterRequest{
		Name: name,
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Printf("The Theater response is %v", res)
}

func getBook(name string, c pb.GreeterClient) {

	log.Println("Getting Book")

	res, err := c.Book(context.Background(), &pb.BookRequest{
		Title:  "Harry Potter",
		Author: "J.K. Rowling",
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Printf("The Book response is %v", res)
}
