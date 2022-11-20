package commonredislistener

type RedisListener[T any] interface {
	Subscribe(func(message T)) (unsubscriber func())
}
