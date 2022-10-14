package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	bootstrapServer string = "localhost:9092"
	groupID         string = "test-go-123"
	topics          []string
	printValue      bool = false
)

func ce(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	p := fmt.Println

	// Arguments to program
	bootstrapServer := flag.String("bootstrap-server", bootstrapServer, "kafka server:port to connect to")
	groupID := flag.String("group-id", groupID, "consumer group-id")
	printValue := flag.Bool("print-value", printValue, "set to 'true' if message values needs to be printed")
	topic := flag.String("topic", "", "topic to consume from")
	flag.Parse()

	if *topic == "" {
		panic("Please provide source topic name.")
	}

	// Initializing kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": *bootstrapServer,
		"group.id":          *groupID,
	})
	ce(err)
	defer consumer.Close()

	topics = []string{*topic}
	err = consumer.SubscribeTopics(topics, nil)
	ce(err)

	// Terminate if there is an inerrupt, eg: Ctrl+C
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			p("Caught interrupt, terminating...", sig)
			run = false
		default:
			event := consumer.Poll(100)
			switch message := event.(type) {
			case *kafka.Message:
				if *printValue {
					p("Partition: ", message.TopicPartition.Partition, "Offset: ", message.TopicPartition.Offset, "Key: ", string(message.Key), ", Value: ", string(message.Value))
				} else {
					p("Partition: ", message.TopicPartition.Partition, "Offset: ", message.TopicPartition.Offset, "Key: ", string(message.Key))
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", message)
				run = false
			}
		}
	}

}
