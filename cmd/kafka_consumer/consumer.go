package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9093"},
		GroupID:     "two",
		GroupTopics: []string{"test", "testdemo", "test1","akashtest"},
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(string(m.Value))
	}
}
