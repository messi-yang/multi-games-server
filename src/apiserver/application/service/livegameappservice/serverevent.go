package livegameappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType              ServerEventType = "ERRORED"
	DimensionUpdatedServerEventType     ServerEventType = "DIMENSION_UPDATED"
	GameJoinedServerEventType           ServerEventType = "GAME_JOINED"
	ItemsUpdatedServerEventType         ServerEventType = "ITEMS_UPDATED"
	RangeObservedServerEventType        ServerEventType = "RANGE_OBSERVED"
	ObservedRangeUpdatedServerEventType ServerEventType = "OBSERVED_RANGE_UPDATED"
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

type GameJoinedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		PlayerId string           `json:"playerId"`
		View     viewmodel.ViewVm `json:"view"`
	} `json:"payload"`
}

type DimensionUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Dimension viewmodel.DimensionVm `json:"dimension"`
	} `json:"payload"`
}

type ItemsUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Items []viewmodel.ItemVm `json:"items"`
	} `json:"payload"`
}

type ObservedRangeUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Range     viewmodel.RangeVm `json:"range"`
		Map       viewmodel.MapVm   `json:"map"`
		UpdatedAt time.Time         `json:"updatedAt"`
	} `json:"payload"`
}

type RangeObservedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Range viewmodel.RangeVm `json:"range"`
		Map   viewmodel.MapVm   `json:"map"`
	} `json:"payload"`
}
