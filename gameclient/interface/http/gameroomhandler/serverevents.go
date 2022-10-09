package gameroomhandler

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
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
	MapSize dto.MapSizeDto `json:"mapSize"`
}
type informationUpdatedEvent struct {
	Type    eventType                      `json:"type"`
	Payload informationUpdatedEventPayload `json:"payload"`
}

func constructInformationUpdatedEvent(mapSize dto.MapSizeDto) *informationUpdatedEvent {
	return &informationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: informationUpdatedEventPayload{
			MapSize: mapSize,
		},
	}
}

type zoomedAreaUpdatedEventPayload struct {
	Area      dto.AreaDto    `json:"area"`
	UnitMap   dto.UnitMapDto `json:"unitMap"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
type zoomedAreaUpdatedEvent struct {
	Type    eventType                     `json:"type"`
	Payload zoomedAreaUpdatedEventPayload `json:"payload"`
}

func constructZoomedAreaUpdatedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) *zoomedAreaUpdatedEvent {
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
	Area    dto.AreaDto    `json:"area"`
	UnitMap dto.UnitMapDto `json:"unitMap"`
}
type areaZoomedEvent struct {
	Type    eventType              `json:"type"`
	Payload areaZoomedEventPayload `json:"payload"`
}

func constructAreaZoomedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) *areaZoomedEvent {
	return &areaZoomedEvent{
		Type: areaZoomedEventType,
		Payload: areaZoomedEventPayload{
			Area:    area,
			UnitMap: unitMap,
		},
	}
}
