package game

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func getBlockTopic(rowIdx int, colIdx int) string {
	return fmt.Sprintf("game-block-%v-%v", colIdx, rowIdx)
}

func WatchBlockUpdateMessages(rowIdx int, colIdx int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{os.Getenv("KAFKA_URL")},
		Topic:       getBlockTopic(rowIdx, colIdx),
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.LastOffset,
	})

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Println(fmt.Sprintf("Top: %v , %s", m.Topic, string(m.Value)))
		// fmt.Println(fmt.Sprintf("Top: %v, Units: %v", m.Topic, string(m.Value)))
		// r.Close()
	}
}

func WriteBlockUpdateMessage(rowIdx int, colIdx int, value []byte) {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(os.Getenv("KAFKA_URL")),
		BatchTimeout:           10 * time.Millisecond,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: getBlockTopic(rowIdx, colIdx),
			Value: value,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
