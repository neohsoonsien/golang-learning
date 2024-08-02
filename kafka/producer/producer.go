package main

import (
	"fmt"
	reader "golang-learning/kafka/utils"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <config-file-path>\n",
			os.Args[0])
		os.Exit(1)
	}
	configFile := os.Args[1]
	conf := reader.ReadConfig(configFile)

	fmt.Printf("The configuration for kafka are %v.\n", conf)

	topic := "purchases"
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for event := range p.Events() {
			fmt.Printf("The event is %v.\n", event)
			switch eventType := event.(type) {
			case *kafka.Message:
				fmt.Printf("THe event type is %v.\n", eventType)
				if eventType.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", eventType.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*eventType.TopicPartition.Topic, string(eventType.Key), string(eventType.Value))
				}
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries", "sport shoes"}

	for n := 0; n < 6; n++ {
		key := users[n]
		data := items[n]
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
	}

	// Wait for all messages to be delivered
	p.Flush(5 * 1000)
	p.Close()
}
