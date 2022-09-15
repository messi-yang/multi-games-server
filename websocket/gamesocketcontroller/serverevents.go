package gamesocketcontroller

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/dto/unitmapdto"
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

func constructInformationUpdatedEvent(mapSizeDto mapsizedto.Dto) *informationUpdatedEvent {
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

func constructZoomedAreaUpdatedEvent(area areadto.Dto, unitMap unitmapdto.Dto) *zoomedAreaUpdatedEvent {
	return &zoomedAreaUpdatedEvent{
		Type: zoomedAreaUpdatedEventType,
		Payload: zoomedAreaUpdatedEventPayload{
			Area:      area,
			UnitMap:   unitMap,
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

func constructAreaZoomedEvent(area areadto.Dto, unitMap unitmapdto.Dto) *areaZoomedEvent {
	return &areaZoomedEvent{
		Type: areaZoomedEventType,
		Payload: areaZoomedEventPayload{
			Area:    area,
			UnitMap: unitMap,
		},
	}
}
