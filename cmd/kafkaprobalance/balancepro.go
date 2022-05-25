package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	c:= kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "testinggroup",
		Balancer: &kafka.Murmur2Balancer{},
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