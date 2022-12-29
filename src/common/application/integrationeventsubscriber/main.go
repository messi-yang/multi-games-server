package integrationeventsubscriber

import "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"

type DeprecatedSubscriber[T event.AppEvent] interface {
	Subscribe(func(event T)) (unsubscriber func())
}

type Subscriber interface {
	Subscribe(channel string, callback func(message []byte)) (unsubscriber func())
}
