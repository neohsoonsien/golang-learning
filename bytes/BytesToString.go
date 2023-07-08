package bytes

import (
	"log"
)

func BytesToString() {
	stringForm := []byte{84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 115, 116, 114, 105, 110, 103}
	log.Printf("The bytes slice is %v", stringForm)
	log.Printf("The string form is %v", string(stringForm))
}