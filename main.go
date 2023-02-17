package main

import (
	"fmt"
	"golang-learning/math"
	// "golang-learning/grpc"
	"golang-learning/marshal"
	"golang-learning/mapping"
	"golang-learning/csv"
	"golang-learning/ioutil"
	"golang-learning/test"
	"golang-learning/mongodb"
	"golang-learning/argument"
)

func main() {
	fmt.Println("This is the main function.")
	fmt.Println(math.Sum(10))

	// grpc.Server()
	marshal.Marshal()
	fmt.Println(mapping.Map("Lily", "C5678")["Lily"])

	csv.Write()

	ioutil.Directory()

	fmt.Println(test.Greet("James"))

	mongodb.Aggregate()
	
	argument.Output()
}
