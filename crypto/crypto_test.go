package crypto

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

func TestDerivePublicPrivateKey(t *testing.T) {
	DerivePublicPrivateKey()
}
