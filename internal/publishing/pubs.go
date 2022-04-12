package publishing

import (
	"bufio"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func Publish(client mqtt.Client, topic string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		token := client.Publish(topic, 2, false, text)
		if text == "exit" {
			break
		}
		token.Wait()
	}

}
