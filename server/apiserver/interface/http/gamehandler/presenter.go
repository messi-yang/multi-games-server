package gamehandler

import (
	"encoding/json"
	"time"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/presenter/presenterdto"
)

type EventType string

const (
	ErrorHappenedEventType             EventType = "ERRORED"
	InformationUpdatedEventType        EventType = "INFORMATION_UPDATED"
	RedisAreaZoomedEventType           EventType = "AREA_ZOOMED"
	RedisZoomedAreaUpdatedEventType    EventType = "ZOOMED_AREA_UPDATED"
	RedisZoomAreaRequestedEventType    EventType = "ZOOM_AREA"
	RedisReviveUnitsRequestedEventType EventType = "REVIVE_UNITS"
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
	Dimension presenterdto.DimensionPresenterDto `json:"dimension"`
}
type InformationUpdatedEvent struct {
	Type    EventType                      `json:"type"`
	Payload InformationUpdatedEventPayload `json:"payload"`
}
type RedisZoomedAreaUpdatedEventPayload struct {
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
	UpdatedAt time.Time                          `json:"updatedAt"`
}
type RedisZoomedAreaUpdatedEvent struct {
	Type    EventType                          `json:"type"`
	Payload RedisZoomedAreaUpdatedEventPayload `json:"payload"`
}

type RedisAreaZoomedEventPayload struct {
	Area      presenterdto.AreaPresenterDto      `json:"area"`
	UnitBlock presenterdto.UnitBlockPresenterDto `json:"unitBlock"`
}
type RedisAreaZoomedEvent struct {
	Type    EventType                   `json:"type"`
	Payload RedisAreaZoomedEventPayload `json:"payload"`
}

type RedisReviveUnitsRequestedEventPayload struct {
	Coordinates []presenterdto.CoordinatePresenterDto `json:"coordinates"`
	ActionedAt  time.Time                             `json:"actionedAt"`
}
type RedisReviveUnitsRequestedEvent struct {
	Type    EventType                             `json:"type"`
	Payload RedisReviveUnitsRequestedEventPayload `json:"payload"`
}

type RedisZoomAreaRequestedEventPayload struct {
	Area       presenterdto.AreaPresenterDto `json:"area"`
	ActionedAt time.Time                     `json:"actionedAt"`
}
type RedisZoomAreaRequestedEvent struct {
	Type    EventType                          `json:"type"`
	Payload RedisZoomAreaRequestedEventPayload `json:"payload"`
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

func (presenter *GameHandlerPresenter) CreateInformationUpdatedEvent(dimension presenterdto.DimensionPresenterDto) InformationUpdatedEvent {
	return InformationUpdatedEvent{
		Type: InformationUpdatedEventType,
		Payload: InformationUpdatedEventPayload{
			Dimension: dimension,
		},
	}
}

func (presenter *GameHandlerPresenter) CreateRedisZoomedAreaUpdatedEvent(area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) RedisZoomedAreaUpdatedEvent {
	return RedisZoomedAreaUpdatedEvent{
		Type: RedisZoomedAreaUpdatedEventType,
		Payload: RedisZoomedAreaUpdatedEventPayload{
			Area:      area,
			UnitBlock: unitBlock,
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *GameHandlerPresenter) CreateRedisAreaZoomedEvent(area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) RedisAreaZoomedEvent {
	return RedisAreaZoomedEvent{
		Type: RedisAreaZoomedEventType,
		Payload: RedisAreaZoomedEventPayload{
			Area:      area,
			UnitBlock: unitBlock,
		},
	}
}

func (presenter *GameHandlerPresenter) ExtractRedisReviveUnitsRequestedEvent(msg []byte) ([]gamecommonmodel.Coordinate, error) {
	var action RedisReviveUnitsRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}
	coordinates, err := presenterdto.ParseCoordinatePresenterDtos(action.Payload.Coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}

func (presenter *GameHandlerPresenter) ExtractRedisZoomAreaRequestedEvent(msg []byte) (gamecommonmodel.Area, error) {
	var action RedisZoomAreaRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return gamecommonmodel.Area{}, err
	}
	area, err := action.Payload.Area.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}

	return area, nil
}
