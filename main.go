package main

import (
	"fmt"

	// "golang-learning/asynchronous"
	// "golang-learning/auth"
	"golang-learning/apache_beam"
	"golang-learning/argument"
	"golang-learning/array"
	"golang-learning/boolean"
	"golang-learning/bytes"
	"golang-learning/csv"
	"golang-learning/enumeration"
	"golang-learning/errors_handler"
	"golang-learning/flagging"
	"golang-learning/float"
	"golang-learning/generic"
	"golang-learning/logger"
	"golang-learning/mapping"
	"golang-learning/marshal"
	"golang-learning/math"
	"golang-learning/method_struct"
	"golang-learning/readfile"

	// "golang-learning/mongodb"
	"golang-learning/outer"
	"golang-learning/pointer"

	// "golang-learning/redis"
	"golang-learning/singleton"
	"golang-learning/slice"
	"golang-learning/string"
	"golang-learning/structs"
	"golang-learning/switch_control"
	"golang-learning/test"
	"golang-learning/timing"

	// "golang-learning/type_assertion"
	"golang-learning/unix_time"
)

func main() {
	array.Array()
	fmt.Println("This is the main function.")
	fmt.Println(math.Sum(10))

	marshal.Marshal()
	marshal.Unmarshal()
	fmt.Println(mapping.MapStruct())
	fmt.Println(mapping.MapStudent("Lily", "C5678")["Lily"])
	fmt.Println(mapping.MapList())

	m := mapping.MapString()
	fmt.Println(mapping.MapPointer(&m))

	csv.Write()

	readfile.ReadFile()
	readfile.ReadDirectory()

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
	structs.StudentInfo()

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

	apache_beam.Hello()

	float.Float()

	timing.Month()

	fmt.Print(mapping.Map2D())

	// mongodb.FindOne()

	fmt.Println(math.AddBigInt(math.BigInt("-200000000"), math.BigInt("1")))
	fmt.Printf("Type: %T Value: %v\n", math.Float64("20"), math.Float64("20"))
	fmt.Printf("big.Float Quotient is %v\n", math.FloatQuotient("2000", "9"))
	fmt.Printf("big.Float Subtract is %v\n", math.FloatSubtract("2000", "10"))

	// type_assertion.TypeCheck("This is a string")

	generic.NonGenericInterface()
	generic.GenericAny()

	// channel.MainChannel("First test of the channel")
}
