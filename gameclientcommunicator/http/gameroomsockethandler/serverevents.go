package gameroomsockethandler

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/presenter/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/presenter/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/gameclientcommunicator/presenter/dto/unitmapdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
)

type eventType string

const (
	errorHappenedEventType      eventType = "ERROR"
	informationUpdatedEventType eventType = "INFORMATION_UPDATED"
	areaZoomedEventType         eventType = "AREA_ZOOMED"
	zoomedAreaUpdatedEventType  eventType = "ZOOMED_AREA_UPDATED"
)

type errorHappenedEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type errorHappenedEvent struct {
	Type    eventType                 `json:"type"`
	Payload errorHappenedEventPayload `json:"payload"`
}

func constructErrorHappenedEvent(clientMessage string) *errorHappenedEvent {
	return &errorHappenedEvent{
		Type: errorHappenedEventType,
		Payload: errorHappenedEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

type informationUpdatedEventPayload struct {
	MapSize mapsizedto.Dto `json:"mapSize"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

func constructInformationUpdatedEvent(mapSize valueobject.MapSize) *informationUpdatedEvent {
	mapSizeDto := mapsizedto.ToDto(mapSize)
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize: mapSizeDto,
		},
	}
}

type zoomedAreaUpdatedEventPayload struct {
	Area      areadto.Dto    `json:"area"`
	UnitMap   unitmapdto.Dto `json:"unitMap"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
type zoomedAreaUpdatedEvent struct {
	Type    eventType                     `json:"type"`
	Payload zoomedAreaUpdatedEventPayload `json:"payload"`
}

func constructZoomedAreaUpdatedEvent(area valueobject.Area, unitMap valueobject.UnitMap) *zoomedAreaUpdatedEvent {
	areaDTO := areadto.ToDTO(area)
	unitMapDTO := unitmapdto.ToDto(&unitMap)
	return &zoomedAreaUpdatedEvent{
		Type: zoomedAreaUpdatedEventType,
		Payload: zoomedAreaUpdatedEventPayload{
			Area:      areaDTO,
			UnitMap:   unitMapDTO,
			UpdatedAt: time.Now(),
		},
	}
}

type areaZoomedEventPayload struct {
	Area    areadto.Dto    `json:"area"`
	UnitMap unitmapdto.Dto `json:"unitMap"`
}
type areaZoomedEvent struct {
	Type    eventType              `json:"type"`
	Payload areaZoomedEventPayload `json:"payload"`
}

func constructAreaZoomedEvent(area valueobject.Area, unitMap valueobject.UnitMap) *areaZoomedEvent {
	areaDTO := areadto.ToDTO(area)
	unitMapDTO := unitmapdto.ToDto(&unitMap)
	return &areaZoomedEvent{
		Type: areaZoomedEventType,
		Payload: areaZoomedEventPayload{
			Area:    areaDTO,
			UnitMap: unitMapDTO,
		},
	}
}
