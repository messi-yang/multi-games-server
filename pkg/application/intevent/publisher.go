package intevent

type Publisher interface {
	Publish(channel string, event Event) error
}
