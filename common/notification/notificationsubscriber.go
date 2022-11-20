package notification

type NotificationSubscriber[T any] interface {
	Subscribe(func(message T)) (unsubscriber func())
}
