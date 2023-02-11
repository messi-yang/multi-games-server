package intevent

type IntEventPublisher interface {
	Publish(channel string, message []byte) error
}
