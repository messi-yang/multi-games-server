package gamehandler

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/livegame/port/dto"
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
	Dimension dto.DimensionDto `json:"dimension"`
}
type InformationUpdatedEvent struct {
	Type    EventType                      `json:"type"`
	Payload InformationUpdatedEventPayload `json:"payload"`
}
type ZoomedAreaUpdatedEventPayload struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
	UpdatedAt time.Time        `json:"updatedAt"`
}
type ZoomedAreaUpdatedEvent struct {
	Type    EventType                     `json:"type"`
	Payload ZoomedAreaUpdatedEventPayload `json:"payload"`
}

type AreaZoomedEventPayload struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
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

type GameHandlerPresenter struct {
}

func NewGameHandlerPresenter() *GameHandlerPresenter {
	return &GameHandlerPresenter{}
}

var gameHandlerPresenter = NewGameHandlerPresenter()

func (presenter *GameHandlerPresenter) ExtractEventType(msg []byte) (EventType, error) {
	var event Event
	err := json.Unmarshal(msg, &event)
	if err != nil {
		return ErrorHappenedEventType, err
	}

	return event.Type, nil
}

func (presenter *GameHandlerPresenter) CreateErroredEvent(clientMessage string) ErroredEvent {
	return ErroredEvent{
		Type: ErrorHappenedEventType,
		Payload: ErroredEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

func (presenter *GameHandlerPresenter) CreateInformationUpdatedEvent(dimension dto.DimensionDto) InformationUpdatedEvent {
	return InformationUpdatedEvent{
		Type: InformationUpdatedEventType,
		Payload: InformationUpdatedEventPayload{
			Dimension: dimension,
		},
	}
}

func (presenter *GameHandlerPresenter) CreateZoomedAreaUpdatedEvent(area dto.AreaDto, unitBlock dto.UnitBlockDto) ZoomedAreaUpdatedEvent {
	return ZoomedAreaUpdatedEvent{
		Type: ZoomedAreaUpdatedEventType,
		Payload: ZoomedAreaUpdatedEventPayload{
			Area:      area,
			UnitBlock: unitBlock,
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *GameHandlerPresenter) CreateAreaZoomedEvent(area dto.AreaDto, unitBlock dto.UnitBlockDto) AreaZoomedEvent {
	return AreaZoomedEvent{
		Type: AreaZoomedEventType,
		Payload: AreaZoomedEventPayload{
			Area:      area,
			UnitBlock: unitBlock,
		},
	}
}

func (presenter *GameHandlerPresenter) ExtractReviveUnitsRequestedEvent(msg []byte) ([]dto.CoordinateDto, error) {
	var action ReviveUnitsRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}
	return action.Payload.Coordinates, nil
}

func (presenter *GameHandlerPresenter) ExtractZoomAreaRequestedEvent(msg []byte) (dto.AreaDto, error) {
	var action ZoomAreaRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return dto.AreaDto{}, err
	}

	return action.Payload.Area, nil
}
