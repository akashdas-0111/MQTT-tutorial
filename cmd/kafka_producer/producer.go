package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/segmentio/kafka-go"
)

func main() {
	c := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9093"),
		Topic:                  "akashtest",
		Balancer:               &kafka.RoundRobin{},
		AllowAutoTopicCreation: false,
	}
	if os.Args[1] == "user" {
		fmt.Println("User Input")
		for {
			fmt.Println("Enter the message")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			msg := scanner.Text()
			err := c.WriteMessages(context.Background(), kafka.Message{Value: []byte(msg)})
			if err != nil {
				fmt.Println("Message not sent: ", err)
				panic("")
			} else {
				fmt.Println("Message sent successfully")
			}

		}
	} else if os.Args[1] == "test" {
		fmt.Println("testing")
		val := "1"
		for {
			err := c.WriteMessages(context.Background(), kafka.Message{Value: []byte(val)})
			fmt.Println(val)
			if err != nil {
				fmt.Println("Message not sent")
			} else {
				fmt.Println("Message sent successfully")
			}
			val = val + "1"
		}
	}
}
