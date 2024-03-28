package goroutine

import (
	"log"
	"time"
)

func routine(from string) {
	for i := 0; i < 3; i++ {
		log.Println(from, ":", i)
	}
}

func MainRoutine() {
	// A goroutine is a lightweight thread of execution.

	// the function call is running synchronously
	routine("Routine 1")

	// function routines are running asynchronously
	go routine("Routine 2")
	go func(msg string) {
		log.Println(msg)
	}("Routine 3")

	time.Sleep(time.Second)
	log.Println("done")
}
