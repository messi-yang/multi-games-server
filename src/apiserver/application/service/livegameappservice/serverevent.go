package livegameappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/itemviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/maprangeviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/mapsizeviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitmapviewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType                 ServerEventType = "ERRORED"
	InformationUpdatedServerEventType      ServerEventType = "INFORMATION_UPDATED"
	ItemsUpdatedServerEventType            ServerEventType = "ITEMS_UPDATED"
	MapRangeObservedServerEventType        ServerEventType = "MAP_RANGE_OBSERVED"
	ObservedMapRangeUpdatedServerEventType ServerEventType = "OBSERVED_MAP_RANGE_UPDATED"
)

type GenericServerEvent struct {
	Type ServerEventType `json:"type"`
}

type ErroredServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		ClientMessage string `json:"clientMessage"`
	} `json:"payload"`
}

type InformationUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		MapSize mapsizeviewmodel.ViewModel `json:"mapSize"`
	} `json:"payload"`
}

type ItemsUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Items []itemviewmodel.ViewModel `json:"items"`
	} `json:"payload"`
}

type ObservedMapRangeUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		MapRange  maprangeviewmodel.ViewModel `json:"mapRange"`
		UnitMap   unitmapviewmodel.ViewModel  `json:"unitMap"`
		UpdatedAt time.Time                   `json:"updatedAt"`
	} `json:"payload"`
}

type MapRangeObservedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		MapRange maprangeviewmodel.ViewModel `json:"mapRange"`
		UnitMap  unitmapviewmodel.ViewModel  `json:"unitMap"`
	} `json:"payload"`
}
