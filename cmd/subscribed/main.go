package main

import (
	"akash-mqtttut/internal/subscription"
	"bufio"
	"context"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/segmentio/kafka-go"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if string(msg.Payload()) == "exit" {
		os.Exit(0)
	}
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	kafkaProducer(string(msg.Payload()))

}
var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}

func kafkaProducer(message string) {
	c := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9093"),
		Topic:                  "akashtest",
		Balancer:               &kafka.RoundRobin{},
		AllowAutoTopicCreation: false,
	}
	err := c.WriteMessages(context.Background(), kafka.Message{Value: []byte(message)})
	fmt.Printf("Sending to kafka: %s\n", message)
	if err != nil {
		fmt.Println("Message not sent: ", err)
		panic("")
	} else {
		fmt.Println("Message sent successfully")
	}
}
func main() {
	refer := mqtt.NewClientOptions()
	refer.AddBroker("localhost:1883")
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
