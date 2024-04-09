package goroutine

import (
	"log"
	"sync"
)

func WaitGroup() {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		log.Print("This can be a placeholder for http/gRPC server 1.")
		waitGroup.Done()
	}()
	waitGroup.Add(1)
	go func() {
		log.Print("This can be a placeholder for http/gRPC server 1.")
		waitGroup.Done()
	}()
	waitGroup.Wait()
}
