package subscription

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Sub(client mqtt.Client,topic string) {
    token := client.Subscribe(topic, 1, nil)
    token.Wait()
  fmt.Printf("Subscribed to topic: %s\n", topic)
}
