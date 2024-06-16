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
func TestReplaceSpaceWithUnderscore(t *testing.T) {
	t.Log("Start the \"ReplaceSpaceWIthUnderscore\" function test.")
	log.Print(ReplaceSpaceWithUnderscore("My Name is Muthu A/L Siva."))
}

func TestRemoveAfterDot(t *testing.T) {
	t.Log("Start the \"RemoveAfterDot\" function test.")
	log.Print(RemoveAfterDot("Filename.jpeg"))
}
