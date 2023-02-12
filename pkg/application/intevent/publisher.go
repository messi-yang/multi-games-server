package intevent

type IntEventPublisher interface {
	Publish(channel string, event IntEvent) error
}
