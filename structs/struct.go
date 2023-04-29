package structs

import (
	"fmt"
)

type student struct {
    name	string
    age  	int
}

func newStudent(name string, age int) (*student) {
	student := student{name: name, age: age}

	return &student
}

func Student() {
	fmt.Println(student{name: "John", age: 29})
	fmt.Println(&student{name: "Steven", age: 25})
	fmt.Println(newStudent("James", 30))
}
