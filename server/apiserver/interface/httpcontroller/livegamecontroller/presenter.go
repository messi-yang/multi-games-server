package livegamecontroller

import (
	"encoding/json"
	"time"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
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
	Dimension commonjsondto.DimensionJsonDto `json:"dimension"`
}
type InformationUpdatedEvent struct {
	Type    EventType                      `json:"type"`
	Payload InformationUpdatedEventPayload `json:"payload"`
}
type ZoomedAreaUpdatedAppEventPayload struct {
	Area      commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
	UpdatedAt time.Time                      `json:"updatedAt"`
}
type ZoomedAreaUpdatedAppEvent struct {
	Type    EventType                        `json:"type"`
	Payload ZoomedAreaUpdatedAppEventPayload `json:"payload"`
}

type AreaZoomedAppEventPayload struct {
	Area      commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
}
type AreaZoomedAppEvent struct {
	Type    EventType                 `json:"type"`
	Payload AreaZoomedAppEventPayload `json:"payload"`
}

type ReviveUnitsRequestedAppEventPayload struct {
	Coordinates []commonjsondto.CoordinateJsonDto `json:"coordinates"`
	ActionedAt  time.Time                         `json:"actionedAt"`
}
type ReviveUnitsRequestedAppEvent struct {
	Type    EventType                           `json:"type"`
	Payload ReviveUnitsRequestedAppEventPayload `json:"payload"`
}

type ZoomAreaRequestedAppEventPayload struct {
	Area       commonjsondto.AreaJsonDto `json:"area"`
	ActionedAt time.Time                 `json:"actionedAt"`
}
type ZoomAreaRequestedAppEvent struct {
	Type    EventType                        `json:"type"`
	Payload ZoomAreaRequestedAppEventPayload `json:"payload"`
}

type Presenter struct {
}

func NewPresenter() *Presenter {
	return &Presenter{}
}

var presenter = NewPresenter()

func (presenter *Presenter) ParseEventType(msg []byte) (EventType, error) {
	var event Event
	err := json.Unmarshal(msg, &event)
	if err != nil {
		return ErrorHappenedEventType, err
	}

	return event.Type, nil
}

func (presenter *Presenter) PresentErroredEvent(clientMessage string) ErroredEvent {
	return ErroredEvent{
		Type: ErrorHappenedEventType,
		Payload: ErroredEventPayload{
			ClientMessage: clientMessage,
		},
	}
}

func (presenter *Presenter) PresentInformationUpdatedEvent(dimension gamecommonmodel.Dimension) InformationUpdatedEvent {
	return InformationUpdatedEvent{
		Type: InformationUpdatedEventType,
		Payload: InformationUpdatedEventPayload{
			Dimension: jsondto.NewDimensionJsonDto(dimension),
		},
	}
}

func (presenter *Presenter) PresentZoomedAreaUpdatedEvent(area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) ZoomedAreaUpdatedAppEvent {
	return ZoomedAreaUpdatedAppEvent{
		Type: ZoomedAreaUpdatedEventType,
		Payload: ZoomedAreaUpdatedAppEventPayload{
			Area:      jsondto.NewAreaJsonDto(area),
			UnitBlock: jsondto.NewUnitBlockJsonDto(unitBlock),
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *Presenter) PresentAreaZoomedEvent(area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) AreaZoomedAppEvent {
	return AreaZoomedAppEvent{
		Type: AreaZoomedEventType,
		Payload: AreaZoomedAppEventPayload{
			Area:      jsondto.NewAreaJsonDto(area),
			UnitBlock: jsondto.NewUnitBlockJsonDto(unitBlock),
		},
	}
}

func (presenter *Presenter) ParseReviveUnitsRequestedEvent(msg []byte) ([]gamecommonmodel.Coordinate, error) {
	var action ReviveUnitsRequestedAppEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return nil, err
	}
	coordinates, err := commonjsondto.ParseCoordinateJsonDtos(action.Payload.Coordinates)
	if err != nil {
		return nil, err
	}

	return coordinates, nil
}

func (presenter *Presenter) ParseZoomAreaRequestedEvent(msg []byte) (gamecommonmodel.Area, error) {
	var action ZoomAreaRequestedAppEvent
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
