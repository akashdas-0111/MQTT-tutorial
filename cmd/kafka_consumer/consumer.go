package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main(){
	 
	connec, _ := kafka.DialLeader(context.Background(),"tcp","localhost:9092","testinggroup",0)
	connec.SetDeadline(time.Now().Add(time.Second*10))
	for{
	message,_ := connec.ReadMessage(1e3)
	fmt.Println(string(message.Value))
	}
	// batch:=connec.ReadBatch(1e3,1e9)
	// bytes:=make([]byte,1e6)
	// for{
	// 	_,err:=batch.Read(bytes)
	// 	if err!=nil{
	// 		break
	// 	}
	// 	fmt.Println(string(bytes))
	// 	fmt.Println()
	// }
}
