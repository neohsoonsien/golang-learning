package regexp

import (
	"log"
	"os"
	"testing"

	"gotest.tools/v3/assert"
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

	// test case for "/" and "\s"
	input := "My Name is Muthu A/L Siva"
	expected := "My_Name_is_Muthu_A_L_Siva"
	assert.Equal(t, ReplaceSpaceWithUnderscore(input), expected)

	// test case for "."
	input = "ref\\2023\\michael.patterson"
	expected = "ref\\2023\\michael_patterson"
	assert.Equal(t, ReplaceSpaceWithUnderscore(input), expected)
}

func TestRemoveAfterDot(t *testing.T) {
	t.Log("Start the \"RemoveAfterDot\" function test.")
	input := "Filename.jpeg"
	expected := "Filename."
	assert.Equal(t, RemoveAfterDot(input), expected)
}

func TestRemoveIncludeDot(t *testing.T) {
	t.Log("Start the \"RemoveIncludeDot\" function test.")
	input := "Image.jpeg"
	expected := "Image"
	assert.Equal(t, RemoveIncludeDot(input), expected)
}
