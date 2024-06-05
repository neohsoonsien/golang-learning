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

func TestStudentInfo(t *testing.T) {
	t.Logf("Start to test the StudentInfo function")

	StudentInfo()
}
