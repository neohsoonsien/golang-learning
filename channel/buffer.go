package channel

import "log"

func Buffer() {
	// channels are `unbuffered`` by default,
	// they will only accept sends (chan <-) if there is a corresponding receiver (<- chan)

	// Buffered channels accept a limited number of values
	// without any corresponding concurrent receiver.
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	log.Println(<-messages)
	log.Println(<-messages)
}
