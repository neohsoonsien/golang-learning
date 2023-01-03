package main

import (
	"fmt"
	"golang-learning/math"
	"golang-learning/grpc"
	"golang-learning/marshal"
)

func main() {
	fmt.Println("This is the main function.")
	fmt.Println(math.Sum(10))

	grpc.Server()
	marshal.Marshal()
}
