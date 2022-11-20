package notification

type NotificationSubscriber[T any] interface {
	Subscribe(func(event T)) (unsubscriber func())
}
