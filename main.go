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
	"golang-learning/decimal128"
	"golang-learning/enumeration"
	"golang-learning/errors_handler"
	"golang-learning/flagging"
	"golang-learning/float"
	"golang-learning/ioutil"
	"golang-learning/logger"
	"golang-learning/mapping"
	"golang-learning/marshal"
	"golang-learning/math"
	"golang-learning/method_struct"

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

	apache_beam.Hello()

	float.Float()

	timing.Month()

	fmt.Print(mapping.Map2D())

	// mongodb.FindOne()

	fmt.Println(math.AddBigInt(math.BigInt("-200000000"), math.BigInt("1")))
	fmt.Printf("Type: %T Value: %v\n", math.Float64("20"), math.Float64("20"))
	fmt.Printf("big.Float Quotient is %v\n", math.FloatQuotient("2000", "9"))
	fmt.Printf("big.Float Subtract is %v\n", math.FloatSubtract("2000", "10"))

	value, _ := decimal128.ConvertToDecimal128(120, 3040)
	fmt.Printf("Value in decimal128, Type: %T Value: %v\n", value, value)
	valueStr := decimal128.ConvertDecimal128ToString(value)
	fmt.Printf("Value in string Type: %T Value: %v\n", valueStr, valueStr)
	decimal, _ := decimal128.ConvertBigIntToDecimal128(math.BigInt("200000000"))
	fmt.Printf("Value in decimal128, Type: %T Value: %v\n", decimal, decimal)

	value, _ = decimal128.ConvertToDecimal128(0, 340000000)
	bigFloat, bigInt, _ := decimal128.ConvertDecimal128ToBig(value)
	fmt.Printf("Value in *big.Float, Type: %T Value: %v\n", bigFloat, bigFloat)
	fmt.Printf("Value in *big.Int, Type: %T Value: %v\n", bigInt, bigInt)

	high, low := value.GetBytes()
	fmt.Printf("High is %v, low is %v\n", high, low)

	integerPart, decimalPart := math.SeparateIntegerDecimal(valueStr)
	fmt.Printf("The integer part is %v, and the decimal part is %v\n", integerPart, decimalPart)
}
