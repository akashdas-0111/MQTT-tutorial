package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"akash-mqtttut/internal/balancer"
	"github.com/segmentio/kafka-go"
)

func main() {
	c := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "kafkatest",
		Balancer: &balancer.Custom{},
	})
	for {
		fmt.Println("Enter the message")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		err := c.WriteMessages(context.Background(), kafka.Message{Value: []byte(msg)})
		if err != nil {
			fmt.Println("Message not sent")
		} else {
			fmt.Println("Message sent successfully")
		}

	}
	// val := "1"
	// for {
	// 	err := c.WriteMessages(context.Background(), kafka.Message{Value: []byte(val)})
	// 	fmt.Println(val)
	// 	if err != nil {
	// 		fmt.Println("Message not sent")
	// 	} else {
	// 		fmt.Println("Message sent successfully")
	// 	}
	// 	val = val + "1"
	// }
}
