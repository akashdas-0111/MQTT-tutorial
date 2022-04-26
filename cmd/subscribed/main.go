package main

import (
	"akash-mqtttut/internal/subscription"
	"bufio"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	if(string(msg.Payload())=="exit"){
		os.Exit(0)
	}
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}

func main() {
	refer := mqtt.NewClientOptions()
	refer.AddBroker("tcp://127.0.0.1:9092")
	refer.SetDefaultPublishHandler(MessagePubHandler)
	refer.OnConnect = ConnectHandler
	refer.OnConnectionLost = ConnectLostHandler
	client := mqtt.NewClient(refer)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Enter the channel name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	channels := scanner.Text()
	subscription.Sub(client, channels)
	client.Disconnect(250)
}
