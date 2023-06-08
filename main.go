package main

import (
	"fmt"
	// "golang-learning/asynchronous"
	// "golang-learning/auth"
	"golang-learning/array"
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
	"golang-learning/pointer"
	// "golang-learning/redis"
	"golang-learning/slice"
	"golang-learning/structs"
	"golang-learning/unix_time"
)

func main() {
	array.Array()
	fmt.Println("This is the main function.")
	fmt.Println(math.Sum(10))

	marshal.Marshal()
	fmt.Println(mapping.MapStruct())
	fmt.Println(mapping.MapStudent("Lily", "C5678")["Lily"])
	fmt.Println(mapping.MapList())
	fmt.Println(mapping.MapString())

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

	pointer.String()
	pointer.Struct()
	pointer.Int()

	structs.Student()

	fmt.Printf("Time now in miliseconds: %v, time now in nanosecond: %v.", unix_time.UnixMilli(), unix_time.UnixNano())
}
