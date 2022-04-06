package broker

import (
	"akash-mqtttut/internal/publishing"
	"akash-mqtttut/internal/subscription"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MessagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
var ConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var ConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}


func Broker(){
	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("testing_mqtt")
	opts.SetUsername("ak")
	opts.SetPassword("Hello")
	opts.SetDefaultPublishHandler(MessagePubHandler)
	opts.OnConnect = ConnectHandler
	opts.OnConnectionLost = ConnectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	
	channels := map[int]string{
		1: "topic/akash",
		2: "topic/quantum",
	}

	fmt.Println(channels)
	var t int
	fmt.Scanf("%d", &t)
	subscription.Sub(client, channels[t])
	publishing.Publish(client, channels[t])
	client.Disconnect(250)
}