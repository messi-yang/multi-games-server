package intevent

type IntEventSubscriber[T IntEvent] interface {
	Subscribe(channel string, handler func(T)) (unsubscriber func())
}
