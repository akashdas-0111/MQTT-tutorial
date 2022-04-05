package main

import (
	"bufio"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
    fmt.Printf("Received message: %s from topic: %s\n",msg.Payload() , msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
    fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
    fmt.Printf("Connect lost: %v", err)
}

func main() {
    var broker = "broker.emqx.io"
    var port = 1883
    opts := mqtt.NewClientOptions()
    opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
    opts.SetClientID("testing_mqtt")
    opts.SetUsername("ak")
    opts.SetPassword("Hello")
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
	channels:= map[int]string{
			1:"topic/testing",
			2:"topic/execute",
	}
	fmt.Println("Enter 1:topic/testing  2:topic/execute")
	var t int
	fmt.Scanf("%d",&t)
    sub(client,channels[t])
    publish(client,channels[t])
    client.Disconnect(250)
}

func publish(client mqtt.Client,topic string) {
	scanner := bufio.NewScanner(os.Stdin)
	for{
		scanner.Scan()
		text :=scanner.Text()
		if(text=="exit"){
			break
		}
        token := client.Publish(topic, 1, false, text)
        token.Wait()
	}

}

func sub(client mqtt.Client,topic string) {
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
  fmt.Printf("Subscribed to topic: %s\n", topic)
}
