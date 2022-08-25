package consumer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func ReadMessageFromTopic() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "tropicana-test", 0)
	conn.SetReadDeadline(time.Now().Add(time.Second * 3))

	batch := conn.ReadBatch(1e3, 1e9)
	message := make([]byte, 1e3)

	for {
		eof, err := batch.Read(message)

		if eof == 0 {
			fmt.Println("Cosumer : All message read successful")
			break
		}

		if err != nil {
			log.Println("Error while reading message from kafka topic", err)
			break
		}
		fmt.Println("Cosumer :", string(message))
	}
}
