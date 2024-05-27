package regexp

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Printf("Begin the test for \"regexp\" package")
	exitVal := m.Run()
	os.Exit(exitVal)
}

// /////////////////////////////////////////////////
// ///////////       Unit Tests      ///////////////
// /////////////////////////////////////////////////
func TestReplace(t *testing.T) {
	t.Log("Start the \"Replace\" function test.")
	log.Print(Replace("My Name is Muthu A/L Siva."))
}
