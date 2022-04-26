package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main(){
	 
	connec, _ := kafka.DialLeader(context.Background(),"tcp","localhost:9092","quickstart-events",0)
	connec.SetDeadline(time.Now().Add(time.Second*10))
	connec.WriteMessages(kafka.Message{Value: []byte("Hello akash this side")})
}