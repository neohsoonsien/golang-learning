package channel

import (
	"log"
	"time"
)

func worker(done chan bool) {
	log.Print("Worker starts...")
	time.Sleep(time.Second)
	log.Println("Worker stops...")

	// send a value to notify that the goroutine is done.
	done <- true
}

func Synchronization() {
	// channels can be used to synchronize execution across goroutines.
	// using a blocking receive force the to wait for a goroutine to finish
	done := make(chan bool, 1)
	go worker(done)

	// block the function from exiting
	// until a notification is received from the worker on the channel.
	<-done
}
