package mqtt

import (
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTPub struct {
	client mqtt.Client
}

// NewMQTTPublisher connects to the broker and returns an MQTTPub
func NewMQTTPublisher(broker string, clientID string) (*MQTTPub, error) {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetConnectTimeout(10 * time.Second)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MQTTPub{client: client}, nil
}

// Publish publishes a payload to a topic
func (m *MQTTPub) Publish(topic string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	token := m.client.Publish(topic, 0, false, data)
	token.Wait()
	return token.Error()
}
