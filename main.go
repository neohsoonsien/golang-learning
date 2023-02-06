package main

import (
	"fmt"
	"golang-learning/math"
	// "golang-learning/grpc"
	"golang-learning/marshal"
	"golang-learning/mapping"
	"golang-learning/csv"
	"golang-learning/ioutil"
)

func main() {
	fmt.Println("This is the main function.")
	fmt.Println(math.Sum(10))

	// grpc.Server()
	marshal.Marshal()
	fmt.Println(mapping.Map("Lily", "C5678")["Lily"])

	csv.Write()

	ioutil.Directory()
}
