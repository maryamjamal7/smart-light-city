package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ListenForStatus(client mqtt.Client) {
	client.Subscribe("city/+/+/status", 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received on %s: %s\n", msg.Topic(), msg.Payload())
	})
}
