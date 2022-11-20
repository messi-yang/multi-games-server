package commonredissubscriber

type RedisSubscriber[T any] interface {
	Subscribe(func(message T)) (unsubscriber func())
}
