package eventbus

type EventBus interface {
	Publish(topic string, payload []byte)
	Subscribe(topic string, handler func(event []byte)) (unsubscriber func())
}
