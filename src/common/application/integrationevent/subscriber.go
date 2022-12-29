package integrationevent

type Subscriber interface {
	Subscribe(channel string, callback func(message []byte)) (unsubscriber func())
}
