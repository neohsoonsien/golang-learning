package mapping

import (
	"fmt"
)

type Student struct {
    Name 		string		`json:"name"`
    Id			string		`json:"id"`
}

func MapStudent(name string, id string) (map[string]*Student) {
	// create empty map
	array := make(map[string]int)

	// insert key/value 
	array["one"] = 1
	array["two"] = 2
	fmt.Println("array: ", array)

	// remove a key/value pairs
	delete(array, "two")
	fmt.Println("array: ", array)

	// return index of a key/value pairs with exist check
	index, exist := array["one"]
    fmt.Println("index: ", index, "value: ", exist)

	// declare studentList
	studentList := make(map[string]*Student, 0)
	studentList[name] = &Student{name, id}

	return studentList
}