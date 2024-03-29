package channel

import "log"

func MainChannel() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	log.Print(msg)
}
