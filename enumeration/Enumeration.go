package enumeration

import (
	"log"
)

type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func Enumeration() {
	state := Running
	log.Printf("state: %v", state)
}
