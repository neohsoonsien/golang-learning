package structs

import (
	"fmt"
)

type Parent struct {
	name string
	age  int
}

type Student struct {
	name   string
	age    int
	parent Parent
}

type Classroom struct {
	name  string
	class string
}

func newStudent(name string, age int) *Student {
	student := Student{name: name, age: age}

	return &student
}

func newClassroom(name string, class string) *Classroom {
	classroom := Classroom{name: name, class: class}

	return &classroom
}

func StudentInfo() {

	var pupil *Student
	pupil = &Student{name: "Jimmy", age: 30}
	fmt.Println(pupil)

	fmt.Println(Student{name: "John", age: 29})
	fmt.Println(&Student{name: "Steven", age: 25})

	var student interface{} = newStudent("James", 30)
	var classroom interface{} = newClassroom("James", "1H")

	studentCheck, ok := student.(*Student)
	fmt.Println(studentCheck, ok)

	classroomCheck, ok := classroom.(*Classroom)
	fmt.Println(classroomCheck, ok)
}
