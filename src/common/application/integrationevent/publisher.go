package integrationevent

type Publisher interface {
	Publish(channel string, message []byte) error
}
