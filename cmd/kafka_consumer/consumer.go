package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	connec, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "testinggrou", 0)
	connec.SetDeadline(time.Now().Add(time.Second * 10))
	for {
		message, _ := connec.ReadMessage(1e3)
		fmt.Println(string(message.Value))
	}
}
