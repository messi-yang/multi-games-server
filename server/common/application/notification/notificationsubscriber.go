package notification

import "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"

type NotificationSubscriber[T event.ApplicationEvent] interface {
	Subscribe(func(event T)) (unsubscriber func())
}
