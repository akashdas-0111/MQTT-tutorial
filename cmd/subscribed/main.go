package main

import (
	"akash-mqtttut/internal/subscription"
	"bufio"
	"fmt"
	"os"
	"context"
	"time"
	"github.com/segmentio/kafka-go"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if(string(msg.Payload())=="exit"){
		os.Exit(0)
	}
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	Kafkaproducer(string(msg.Payload()))
}
var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}
func Kafkaproducer(message string){
	connec, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "testinggroup", 1)
	connec.SetDeadline(time.Now().Add(time.Second * 10))
	connec.WriteMessages(kafka.Message{Value: []byte(message)})
	fmt.Printf("Published message to Kafka: %s to topic quickstart-events\n",message)
}

func main() {
	refer := mqtt.NewClientOptions()
	refer.AddBroker("tcp://127.0.0.1:1883")
	refer.SetDefaultPublishHandler(MessagePubHandler)
	refer.OnConnect = ConnectHandler
	refer.OnConnectionLost = ConnectLostHandler
	client := mqtt.NewClient(refer)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Enter the topic name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	channels := scanner.Text()
	subscription.Sub(client, channels)
	client.Disconnect(250)
}
