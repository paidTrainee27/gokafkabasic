package main

import (
	"flag"
	"fmt"
	"gokafkabasic/consumer"
	"gokafkabasic/producer"
	"time"
)

var (
	kafkaMessage = flag.String("message", "dummy", "No message was passed")
)

func main() {
	flag.Parse()

	go producer.WriteMessage(kafkaMessage)
	go consumer.ReadMessageFromTopic()
	time.Sleep(time.Second * 5)
	fmt.Println("Exiting")
}
