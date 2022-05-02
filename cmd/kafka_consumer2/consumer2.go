package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main(){
	conf:=kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic: "quickstart-events",

	}
	reader:=kafka.NewReader(conf)
	for{
		m,err:=reader.ReadMessage(context.Background())
		if err!=nil{
			fmt.Println("Error",err)
		}
		fmt.Println(string(m.Value))
		
	}
}