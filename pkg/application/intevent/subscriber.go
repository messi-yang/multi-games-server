package intevent

type IntEventSubscriber interface {
	Subscribe(channel string, handler func([]byte)) (unsubscriber func())
}
