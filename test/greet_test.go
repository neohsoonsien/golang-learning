package test

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

func TestGreet(t *testing.T) {
	name := "James"
	t.Logf("The function Hello returns %v", Greet(name))
}

func TestA(t *testing.T) {
	t.Log("TestA running")
}

func TestB(t *testing.T) {
	t.Log("TestB running")
}
