package notification

import "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"

type NotificationPublisher interface {
	Publish(channel string, event event.ApplicationEvent) error
}
