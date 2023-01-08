package intgrevent

type IntgrEventSubscriber interface {
	Subscribe(channel string, handler func([]byte)) (unsubscriber func())
}
