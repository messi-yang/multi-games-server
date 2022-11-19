package notification

type NotificationPublisher interface {
	Publish(channel string, jsonMessage any) error
}
