package readfile

import (
	"log"
	"os"
)

func ReadFile() {

	content, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	log.Printf("The \"content\" in []byte: %v", content)
	log.Printf("The \"content\" in string: %v", content)
}
