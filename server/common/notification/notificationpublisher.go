package notification

type NotificationPublisher interface {
	Publish(channel string, event any) error
}
