package ports

type MQTTPublisher interface {
	Publish(topic string, payload interface{}) error
}
