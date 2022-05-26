package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9093", "test1", 0)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	//kafka.Dial("tcp", "localhost:9093","test")
	partitions, err := conn.ReadPartitions()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(partitions))
}
