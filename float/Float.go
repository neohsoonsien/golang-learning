package float

import (
	"log"
	"strconv"
)

func Float() {
	str := "3.14159265"
	// value will be in float64 type, even though ParseFloat parameter is 32 bit
	if value, err := strconv.ParseFloat(str, 32); err == nil {
		log.Print(value) // 3.1415927410125732
		// convert back to string
		str = strconv.FormatFloat(value, 'E', -1, 32)
		log.Print(str)
	}
	if value, err := strconv.ParseFloat(str, 64); err == nil {
		log.Print(value) // 3.14159265
		// convert back to string
		str = strconv.FormatFloat(value, 'E', -1, 64)
		log.Print(str)
	}
}
