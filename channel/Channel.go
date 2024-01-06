package channel

import "log"

func subChannelOne(channel chan string) {

	log.Printf("The message is %v", <-channel)
	close(channel)
}

func MainChannel(message string) {
	log.Print("Start of the main channel")
	channel := make(chan string)
	subChannelOne(channel)

	log.Print("Send the message to the channel")
	channel <- message
	log.Print("End of the main channel")
}
