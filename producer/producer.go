package producer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func WriteMessage(message *string) {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "tropicana-test", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{
		Value: []byte(*message),
	})
}
