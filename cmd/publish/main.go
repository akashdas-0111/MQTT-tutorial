package main

import (
	"akash-mqtttut/internal/publishing"
	"bufio"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)


var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}

func main() {
	refer := mqtt.NewClientOptions()
	refer.AddBroker("tcp://127.0.0.1:9092")
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
	publishing.Publish(client, channels)
	client.Disconnect(250)
}
