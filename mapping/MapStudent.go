package mapping

import (
	"fmt"
)

type Teacher struct {
	Name     string              `json:"Name"`
	Subjects map[string][]string `json:"Subjects"`
}

type Student struct {
	Name string `json:"Name"`
	Id   string `json:"Id"`
}

func MapStudent(name string, id string) map[string]*Student {
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
	studentList[name] = &Student{
		Name: name,
		Id:   id,
	}

	// the object component is a map
	classes := []string{"1H", "2B", "2A"}
	subjectsList := make(map[string][]string)
	subjectsList["History"] = classes
	teacher := Teacher{
		Name:     "Jenny",
		Subjects: subjectsList,
	}
	for subject, classes := range teacher.Subjects {
		fmt.Printf("Staff name: %v, subject taught: %v, classes: %v\n", teacher.Name, subject, classes)
	}

	return studentList
}
