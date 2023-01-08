package intgrevent

type IntgrEventPublisher interface {
	Publish(channel string, message []byte) error
}
