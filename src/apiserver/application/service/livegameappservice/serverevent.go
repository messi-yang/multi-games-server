package livegameappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
)

type ServerEventType string

const (
	ErroredServerEventType              ServerEventType = "ERRORED"
	InformationUpdatedServerEventType   ServerEventType = "INFORMATION_UPDATED"
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

type InformationUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		MapSize viewmodel.MapSize `json:"mapSize"`
	} `json:"payload"`
}

type ItemsUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Items []viewmodel.Item `json:"items"`
	} `json:"payload"`
}

type ObservedRangeUpdatedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Range     viewmodel.Range   `json:"range"`
		UnitMap   viewmodel.UnitMap `json:"unitMap"`
		UpdatedAt time.Time         `json:"updatedAt"`
	} `json:"payload"`
}

type RangeObservedServerEvent struct {
	Type    ServerEventType `json:"type"`
	Payload struct {
		Range   viewmodel.Range   `json:"range"`
		UnitMap viewmodel.UnitMap `json:"unitMap"`
	} `json:"payload"`
}
