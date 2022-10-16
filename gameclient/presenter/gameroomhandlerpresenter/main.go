package gameroomhandlerpresenter

import (
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
)

type EventType string

const (
	errorHappenedEventType      EventType = "ERRORED"
	informationUpdatedEventType EventType = "INFORMATION_UPDATED"
	areaZoomedEventType         EventType = "AREA_ZOOMED"
	zoomedAreaUpdatedEventType  EventType = "ZOOMED_AREA_UPDATED"
)

type ErroredEventPayload struct {
	ClientMessage string `json:"clientMessage"`
}
type ErroredEvent struct {
	Type    EventType           `json:"type"`
	Payload ErroredEventPayload `json:"payload"`
}

type InformationUpdatedEventPayload struct {
	MapSize dto.MapSizeDto `json:"mapSize"`
}
type InformationUpdatedEvent struct {
	Type    EventType                      `json:"type"`
	Payload InformationUpdatedEventPayload `json:"payload"`
}
type ZoomedAreaUpdatedEventPayload struct {
	Area      dto.AreaDto    `json:"area"`
	UnitMap   dto.UnitMapDto `json:"unitMap"`
	UpdatedAt time.Time      `json:"updatedAt"`
}
type ZoomedAreaUpdatedEvent struct {
	Type    EventType                     `json:"type"`
	Payload ZoomedAreaUpdatedEventPayload `json:"payload"`
}

type AreaZoomedEventPayload struct {
	Area    dto.AreaDto    `json:"area"`
	UnitMap dto.UnitMapDto `json:"unitMap"`
}
type AreaZoomedEvent struct {
	Type    EventType              `json:"type"`
	Payload AreaZoomedEventPayload `json:"payload"`
}

type GameRoomHandlerPresenter interface {
	CreateErroredEvent(clientMessage string) ErroredEvent
	CreateInformationUpdatedEvent(mapSize dto.MapSizeDto) InformationUpdatedEvent
	CreateZoomedAreaUpdatedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) ZoomedAreaUpdatedEvent
	CreateAreaZoomedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) AreaZoomedEvent
}

type gameRoomHandlerPresenterImplement struct {
}

func NewGameRoomHandlerPresenter() GameRoomHandlerPresenter {
	return &gameRoomHandlerPresenterImplement{}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateErroredEvent(clientMessage string) ErroredEvent {
	return ErroredEvent{
		Type: errorHappenedEventType,
		Payload: ErroredEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateInformationUpdatedEvent(mapSize dto.MapSizeDto) InformationUpdatedEvent {
	return InformationUpdatedEvent{
		Type: informationUpdatedEventType,
		Payload: InformationUpdatedEventPayload{
			MapSize: mapSize,
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateZoomedAreaUpdatedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) ZoomedAreaUpdatedEvent {
	return ZoomedAreaUpdatedEvent{
		Type: zoomedAreaUpdatedEventType,
		Payload: ZoomedAreaUpdatedEventPayload{
			Area:      area,
			UnitMap:   unitMap,
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateAreaZoomedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) AreaZoomedEvent {
	return AreaZoomedEvent{
		Type: areaZoomedEventType,
		Payload: AreaZoomedEventPayload{
			Area:    area,
			UnitMap: unitMap,
		},
	}
}
