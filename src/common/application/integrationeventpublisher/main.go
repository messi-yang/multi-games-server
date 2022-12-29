package integrationeventpublisher

import "github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/event"

type Publisher interface {
	Publish(channel string, event event.AppEvent) error
}
