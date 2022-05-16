package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)
type Custom struct {
	// contains filtered or unexported fields
}
func (lb *Custom) Balance(msg kafka.Message, partitions ...int)int{
	return 0
}
func main() {
	var part int
	fmt.Println("Enter partition number")
	fmt.Scanf("%d",part)
	connec, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "testinggroup",part)
	connec.SetDeadline(time.Now().Add(time.Second * 10))
	fmt.Println("Enter the message")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()
	connec.WriteMessages(kafka.Message{Value: []byte(msg)})
	
}
