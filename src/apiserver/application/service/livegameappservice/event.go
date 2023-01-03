package livegameappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/gamemapviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/itemviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/maprangeviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/mapsizeviewmodel"
)

type EventType string

const (
	ErrorHappenedEventType           EventType = "ERRORED"
	InformationUpdatedEventType      EventType = "INFORMATION_UPDATED"
	ItemsUpdatedEventType            EventType = "ITEMS_UPDATED"
	MapRangeObservedEventType        EventType = "MAP_RANGE_OBSERVED"
	ObservedMapRangeUpdatedEventType EventType = "OBSERVED_MAP_RANGE_UPDATED"
)

type GenericEvent struct {
	Type EventType `json:"type"`
}

type ErroredEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		ClientMessage string `json:"clientMessage"`
	} `json:"payload"`
}

type InformationUpdatedEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		MapSize mapsizeviewmodel.ViewModel `json:"mapSize"`
	} `json:"payload"`
}

type ItemsUpdatedEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		Items []itemviewmodel.ViewModel `json:"items"`
	} `json:"payload"`
}

type ObservedMapRangeUpdatedEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		MapRange  maprangeviewmodel.ViewModel `json:"mapRange"`
		GameMap   gamemapviewmodel.ViewModel  `json:"gameMap"`
		UpdatedAt time.Time                   `json:"updatedAt"`
	} `json:"payload"`
}

type MapRangeObservedEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		MapRange maprangeviewmodel.ViewModel `json:"mapRange"`
		GameMap  gamemapviewmodel.ViewModel  `json:"gameMap"`
	} `json:"payload"`
}
