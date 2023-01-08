package livegameappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType               ServerEventType = "ERRORED"
	InformationUpdatedServerEventType    ServerEventType = "INFORMATION_UPDATED"
	ItemsUpdatedServerEventType          ServerEventType = "ITEMS_UPDATED"
	ExtentObservedServerEventType        ServerEventType = "MAP_RANGE_OBSERVED"
	ObservedExtentUpdatedServerEventType ServerEventType = "OBSERVED_MAP_RANGE_UPDATED"
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
		MapSize viewmodel.MapSizeViewModel `json:"mapSize"`
	} `json:"payload"`
}

type ItemsUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Items []viewmodel.ItemViewModel `json:"items"`
	} `json:"payload"`
}

type ObservedExtentUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Extent    viewmodel.ExtentViewModel  `json:"extent"`
		UnitMap   viewmodel.UnitMapViewModel `json:"unitMap"`
		UpdatedAt time.Time                  `json:"updatedAt"`
	} `json:"payload"`
}

type ExtentObservedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Extent  viewmodel.ExtentViewModel  `json:"extent"`
		UnitMap viewmodel.UnitMapViewModel `json:"unitMap"`
	} `json:"payload"`
}
