package subscription

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 2, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s\n", topic)
	time.Sleep(250 * time.Second)
}
