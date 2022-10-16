package gameroomhandlerpresenter

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
)

type EventType string

const (
	ErrorHappenedEventType        EventType = "ERRORED"
	InformationUpdatedEventType   EventType = "INFORMATION_UPDATED"
	AreaZoomedEventType           EventType = "AREA_ZOOMED"
	ZoomedAreaUpdatedEventType    EventType = "ZOOMED_AREA_UPDATED"
	ZoomAreaRequestedEventType    EventType = "ZOOM_AREA"
	ReviveUnitsRequestedEventType EventType = "REVIVE_UNITS"
)

type Event struct {
	Type    EventType   `json:"type"`
	Payload interface{} `json:"payload"`
}

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

type ReviveUnitsRequestedEventPayload struct {
	Coordinates []dto.CoordinateDto `json:"coordinates"`
	ActionedAt  time.Time           `json:"actionedAt"`
}
type ReviveUnitsRequestedEvent struct {
	Type    EventType                        `json:"type"`
	Payload ReviveUnitsRequestedEventPayload `json:"payload"`
}

type ZoomAreaRequestedEventPayload struct {
	Area       dto.AreaDto `json:"area"`
	ActionedAt time.Time   `json:"actionedAt"`
}
type ZoomAreaRequestedEvent struct {
	Type    EventType                     `json:"type"`
	Payload ZoomAreaRequestedEventPayload `json:"payload"`
}

type GameRoomHandlerPresenter interface {
	CreateErroredEvent(clientMessage string) ErroredEvent
	CreateInformationUpdatedEvent(mapSize dto.MapSizeDto) InformationUpdatedEvent
	CreateZoomedAreaUpdatedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) ZoomedAreaUpdatedEvent
	CreateAreaZoomedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) AreaZoomedEvent
	ExtractEventType(msg []byte) (EventType, error)
	ExtractReviveUnitsRequestedEvent(msg []byte) ([]valueobject.Coordinate, error)
	ExtractZoomAreaRequestedEvent(msg []byte) (valueobject.Area, error)
}

type gameRoomHandlerPresenterImplement struct {
}

func NewGameRoomHandlerPresenter() GameRoomHandlerPresenter {
	return &gameRoomHandlerPresenterImplement{}
}

func (presenter *gameRoomHandlerPresenterImplement) ExtractEventType(msg []byte) (EventType, error) {
	var event Event
	err := json.Unmarshal(msg, &event)
	if err != nil {
		return ErrorHappenedEventType, err
	}

	return event.Type, nil
}

func (presenter *gameRoomHandlerPresenterImplement) CreateErroredEvent(clientMessage string) ErroredEvent {
	return ErroredEvent{
		Type: ErrorHappenedEventType,
		Payload: ErroredEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateInformationUpdatedEvent(mapSize dto.MapSizeDto) InformationUpdatedEvent {
	return InformationUpdatedEvent{
		Type: InformationUpdatedEventType,
		Payload: InformationUpdatedEventPayload{
			MapSize: mapSize,
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateZoomedAreaUpdatedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) ZoomedAreaUpdatedEvent {
	return ZoomedAreaUpdatedEvent{
		Type: ZoomedAreaUpdatedEventType,
		Payload: ZoomedAreaUpdatedEventPayload{
			Area:      area,
			UnitMap:   unitMap,
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) CreateAreaZoomedEvent(area dto.AreaDto, unitMap dto.UnitMapDto) AreaZoomedEvent {
	return AreaZoomedEvent{
		Type: AreaZoomedEventType,
		Payload: AreaZoomedEventPayload{
			Area:    area,
			UnitMap: unitMap,
		},
	}
}

func (presenter *gameRoomHandlerPresenterImplement) ExtractReviveUnitsRequestedEvent(msg []byte) ([]valueobject.Coordinate, error) {
	var action ReviveUnitsRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}

	coordinates, err := dto.ParseCoordinateDtos(action.Payload.Coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}

func (presenter *gameRoomHandlerPresenterImplement) ExtractZoomAreaRequestedEvent(msg []byte) (valueobject.Area, error) {
	var action ZoomAreaRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return valueobject.Area{}, err
	}

	area, err := action.Payload.Area.ToValueObject()
	if err != nil {
		return valueobject.Area{}, err
	}

	return area, nil
}
