package channel

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

func TestMainChannel(t *testing.T) {
	MainChannel()
}

func TestBuffer(t *testing.T) {
	Buffer()
}

func TestClose(t *testing.T) {
	Close()
}

func TestSynchronization(t *testing.T) {
	Synchronization()
}

func TestTimeout(*testing.T) {
	Timeout()
}
