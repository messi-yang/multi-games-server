package eventbus

type IntegrationEventBus[T any] interface {
	Publish(topic string, payload T)
}
