package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "two",
		Topic:   "testinggroup",
		// Partition: 1,

	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(string(m.Value))
		time.Sleep(5*time.Second)
	}
}
