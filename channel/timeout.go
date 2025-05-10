package channel

import (
	"context"
	"fmt"
	"time"
)

func Timeout() {
	// define the context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// define the channel
	result := make(chan string, 1)

	// go-routine
	go func() {
		time.Sleep(3 * time.Second)
		answer := "Hello"

		result <- answer
	}()

	// output the results
	select {
	case r := <-result:
		fmt.Println("answer is:", r)
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}
