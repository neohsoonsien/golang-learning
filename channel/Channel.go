package channel

func ChannelOne() string {
	messages := make(chan string)

	ChannelTwo(messages)

	return <-messages
}

func ChannelTwo(msg chan string) {
	msg <- "message from ChannelTwo"
}
