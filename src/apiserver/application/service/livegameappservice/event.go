package livegameappservice

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/dimensionviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitblockviewmodel"
)

type EventType string

const (
	ErrorHappenedEventType      EventType = "ERRORED"
	InformationUpdatedEventType EventType = "INFORMATION_UPDATED"
	AreaZoomedEventType         EventType = "AREA_ZOOMED"
	ZoomedAreaUpdatedEventType  EventType = "ZOOMED_AREA_UPDATED"
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
		Dimension dimensionviewmodel.ViewModel `json:"dimension"`
	} `json:"payload"`
}

type ZoomedAreaUpdatedEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		Area      areaviewmodel.ViewModel      `json:"area"`
		UnitBlock unitblockviewmodel.ViewModel `json:"unitBlock"`
		UpdatedAt time.Time                    `json:"updatedAt"`
	} `json:"payload"`
}

type AreaZoomedEvent struct {
	Type    EventType `json:"type"`
	Payload struct {
		Area      areaviewmodel.ViewModel      `json:"area"`
		UnitBlock unitblockviewmodel.ViewModel `json:"unitBlock"`
	} `json:"payload"`
}
