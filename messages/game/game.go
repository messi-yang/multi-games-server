package game

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

func getBlockTopic(rowIdx int, colIdx int) string {
	return fmt.Sprintf("game-block-update-%v-%v", colIdx, rowIdx)
}

func ListTopics() {
	conn, err := kafka.Dial("tcp", os.Getenv("KAFKA_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}
	for k := range m {
		fmt.Println(k)
	}
}

func WatchBlockUpdateMessages(rowIdx int, colIdx int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_URL")},
		Topic:    getBlockTopic(rowIdx, colIdx),
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	ctx := context.Background()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		messageTime := m.Time.Add(200 * time.Millisecond).UTC()
		currentTime := time.Now().UTC()
		isNewData := messageTime.After(currentTime)

		if isNewData {
			fmt.Println(fmt.Sprintf("Top: %v", m.Topic))
		}
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
		log.Printf("failed to write message: %s", err.Error())
	}

	if err := w.Close(); err != nil {
		log.Printf("failed to close writer: %s", err.Error())
	}
}

func CreateGameRelatedTopics() {
	scale, _ := strconv.Atoi(os.Getenv("GAME_MAP_SCALE"))
	conn, err := kafka.Dial("tcp", os.Getenv("KAFKA_URL"))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := make([]kafka.TopicConfig, 0)

	for rowIdx := 0; rowIdx < scale; rowIdx += 1 {
		for colIdx := 0; colIdx < scale; colIdx += 1 {
			topicConfigs = append(topicConfigs, kafka.TopicConfig{
				Topic:             getBlockTopic(rowIdx, colIdx),
				NumPartitions:     3,
				ReplicationFactor: 1,
			})
		}
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
