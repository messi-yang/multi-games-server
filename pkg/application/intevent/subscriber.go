package intevent

type Subscriber[T Event] interface {
	Subscribe(channel string, handler func(T)) (unsubscriber func())
}
