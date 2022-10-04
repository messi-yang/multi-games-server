package eventbus

type IntegrationEventBus interface {
	Publish(topic string, payload []byte)
	Subscribe(topic string, handler func(event []byte)) (unsubscriber func())
}
