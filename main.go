package main

import (
	"fmt"
	// "golang-learning/asynchronous"
	// "golang-learning/auth"
	"golang-learning/array"
	"golang-learning/boolean"
	"golang-learning/bytes"
	"golang-learning/csv"
	"golang-learning/enumeration"
	"golang-learning/errors_handler"
	"golang-learning/flagging"
	"golang-learning/ioutil"
	"golang-learning/mapping"
	"golang-learning/marshal"
	"golang-learning/math"
	"golang-learning/method_struct"
	"golang-learning/test"

	// "golang-learning/mongodb"
	"golang-learning/argument"
	"golang-learning/logger"
	"golang-learning/outer"
	"golang-learning/pointer"

	// "golang-learning/redis"
	"golang-learning/singleton"
	"golang-learning/slice"
	"golang-learning/string"
	"golang-learning/structs"
	"golang-learning/switch_control"
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

	m := mapping.MapString()
	fmt.Println(mapping.MapPointer(&m))

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

	// asynchronous.Asynchronous()

	// mongodb.Group()

	// mongodb.Match()

	// mongodb.Sort()

	// mongodb.FindOne_UpdateOne()

	outer.Exit()

	boolean.Boolean()

	flagging.DefineColor()

	method_struct.Method()

	pointer.String()
	pointer.Struct()
	pointer.Int()

	string.Prefix("This is a string", "This")
	structs.Student()

	fmt.Printf("Time now in miliseconds: %v, time now in nanosecond: %v.\n", unix_time.UnixMilli(), unix_time.UnixNano())

	// mongodb.Aggregate2()

	// mongodb.BulkWrite()

	switch_control.Switch("C")

	bytes.StringToBytes()
	bytes.BytesToString()

	errors_handler.Handler()

	s1 := singleton.GetInstance("first")
	s2 := singleton.GetInstance("second")
	println("s1.GetName()", s1.GetName())
	println("s2.GetName()", s2.GetName())
	println("s1 == s2", s1 == s2)

	enumeration.Enumeration()
}
