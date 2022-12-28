package notification

import "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"

type NotificationPublisher interface {
	Publish(channel string, event event.AppEvent) error
}
