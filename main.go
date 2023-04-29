package main

import (
	"fmt"
	// "golang-learning/asynchronous"
	// "golang-learning/auth"
	"golang-learning/boolean"
	"golang-learning/math"
	"golang-learning/marshal"
	"golang-learning/mapping"
	"golang-learning/method_struct"
	"golang-learning/csv"
	"golang-learning/flagging"
	"golang-learning/ioutil"
	"golang-learning/test"
	// "golang-learning/mongodb"
	"golang-learning/outer"
	"golang-learning/argument"
	"golang-learning/logger"
	// "golang-learning/redis"
	"golang-learning/slice"
	"golang-learning/structs"
)

func main() {
	fmt.Println("This is the main function.")
	fmt.Println(math.Sum(10))

	marshal.Marshal()
	fmt.Println(mapping.MapStruct())
	fmt.Println(mapping.MapStudent("Lily", "C5678")["Lily"])
	fmt.Println(mapping.MapList())

	csv.Write()

	ioutil.Directory()

	fmt.Println(test.Greet("James"))

	argument.Output()

	logger.Logger()

	// mongodb.Find()

	// mongodb.Insert()

	// mongodb.And()

	// mongodb.Aggregate()
	
	// mongodb.ElementMatch()

	// mongodb.UpdateMany()

	// redis.GetStudent()

	slice.Slice()

	// auth.Auth()

	// auth.GrpcAuth()

	// asynchronous.Asynchronous()

	// mongodb.Group()

	// mongodb.Match()

	// mongodb.Sort()

	outer.Exit()

	boolean.Boolean()

	flagging.DefineColor()

	method_struct.Method()

	structs.Student()
}
