package eventbus

type IntegrationEventBus[T any] interface {
	Publish(topic string, payload T)
	Subscribe(topic string, handler func(event T)) (unsubscriber func())
}
