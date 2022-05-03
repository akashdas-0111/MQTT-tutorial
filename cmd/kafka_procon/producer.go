package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	connec, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "testinggroup",0)
	connec.SetDeadline(time.Now().Add(time.Second * 10))
	fmt.Println("Enter the message")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()
	connec.WriteMessages(kafka.Message{Value: []byte(msg)})
	
}
