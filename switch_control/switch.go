package switch_control

import (
	"log"
)

func Switch(option string) {
	switch option {
	case "A":
		log.Print(option)
	case "B":
		log.Print(option)
	case "C", "D":
		log.Print("C or D")
	default:
		log.Print("default")
	}
}