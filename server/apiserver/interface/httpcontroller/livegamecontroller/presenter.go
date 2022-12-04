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
type ZoomedAreaUpdatedApplicationEventPayload struct {
	Area      commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
	UpdatedAt time.Time                      `json:"updatedAt"`
}
type ZoomedAreaUpdatedApplicationEvent struct {
	Type    EventType                                `json:"type"`
	Payload ZoomedAreaUpdatedApplicationEventPayload `json:"payload"`
}

type AreaZoomedApplicationEventPayload struct {
	Area      commonjsondto.AreaJsonDto      `json:"area"`
	UnitBlock commonjsondto.UnitBlockJsonDto `json:"unitBlock"`
}
type AreaZoomedApplicationEvent struct {
	Type    EventType                         `json:"type"`
	Payload AreaZoomedApplicationEventPayload `json:"payload"`
}

type ReviveUnitsRequestedApplicationEventPayload struct {
	Coordinates []commonjsondto.CoordinateJsonDto `json:"coordinates"`
	ActionedAt  time.Time                         `json:"actionedAt"`
}
type ReviveUnitsRequestedApplicationEvent struct {
	Type    EventType                                   `json:"type"`
	Payload ReviveUnitsRequestedApplicationEventPayload `json:"payload"`
}

type ZoomAreaRequestedApplicationEventPayload struct {
	Area       commonjsondto.AreaJsonDto `json:"area"`
	ActionedAt time.Time                 `json:"actionedAt"`
}
type ZoomAreaRequestedApplicationEvent struct {
	Type    EventType                                `json:"type"`
	Payload ZoomAreaRequestedApplicationEventPayload `json:"payload"`
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

func (presenter *Presenter) PresentZoomedAreaUpdatedEvent(area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) ZoomedAreaUpdatedApplicationEvent {
	return ZoomedAreaUpdatedApplicationEvent{
		Type: ZoomedAreaUpdatedEventType,
		Payload: ZoomedAreaUpdatedApplicationEventPayload{
			Area:      jsondto.NewAreaJsonDto(area),
			UnitBlock: jsondto.NewUnitBlockJsonDto(unitBlock),
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *Presenter) PresentAreaZoomedEvent(area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) AreaZoomedApplicationEvent {
	return AreaZoomedApplicationEvent{
		Type: AreaZoomedEventType,
		Payload: AreaZoomedApplicationEventPayload{
			Area:      jsondto.NewAreaJsonDto(area),
			UnitBlock: jsondto.NewUnitBlockJsonDto(unitBlock),
		},
	}
}

func (presenter *Presenter) ParseReviveUnitsRequestedEvent(msg []byte) ([]gamecommonmodel.Coordinate, error) {
	var action ReviveUnitsRequestedApplicationEvent
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
	var action ZoomAreaRequestedApplicationEvent
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
