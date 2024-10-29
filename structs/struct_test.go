package structs

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Print("BEFORE the tests!")
	exitVal := m.Run()
	log.Print("AFTER the tests!")

	os.Exit(exitVal)
}

func TestNewStudent(t *testing.T) {
	t.Logf("Start to test the newStudent function")

	student := newStudent("Tim", 17)

	if student == nil {
		t.Errorf("Failed to create the pointer of 'Student'")
	}

	if student.parent.name != "" {
		t.Errorf("The 'parent.name' of 'student' should be \"\".")
	}

	if student.parent.age != 0 {
		t.Errorf("The 'parent.age' of 'student' should be zero.")
	}
}

func TestStudentInfo(t *testing.T) {
	t.Logf("Start to test the StudentInfo function")

	StudentInfo()
}
