package livegamecontroller

import (
	"encoding/json"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/dimensionviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitblockviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
)

type EventType string

const (
	ErrorHappenedEventType      EventType = "ERRORED"
	InformationUpdatedEventType EventType = "INFORMATION_UPDATED"
	AreaZoomedEventType         EventType = "AREA_ZOOMED"
	ZoomedAreaUpdatedEventType  EventType = "ZOOMED_AREA_UPDATED"
	ZoomAreaEventType           EventType = "ZOOM_AREA"
	BuildItemEventType          EventType = "BUILD_ITEM"
	DestroyItemEventType        EventType = "DESTROY_ITEM"
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
	Dimension dimensionviewmodel.ViewModel `json:"dimension"`
}
type InformationUpdatedEvent struct {
	Type    EventType                      `json:"type"`
	Payload InformationUpdatedEventPayload `json:"payload"`
}
type ZoomedAreaUpdatedEventPayload struct {
	Area      areaviewmodel.ViewModel      `json:"area"`
	UnitBlock unitblockviewmodel.ViewModel `json:"unitBlock"`
	UpdatedAt time.Time                    `json:"updatedAt"`
}
type ZoomedAreaUpdatedEvent struct {
	Type    EventType                     `json:"type"`
	Payload ZoomedAreaUpdatedEventPayload `json:"payload"`
}

type AreaZoomedEventPayload struct {
	Area      areaviewmodel.ViewModel      `json:"area"`
	UnitBlock unitblockviewmodel.ViewModel `json:"unitBlock"`
}
type AreaZoomedEvent struct {
	Type    EventType              `json:"type"`
	Payload AreaZoomedEventPayload `json:"payload"`
}

type BuildItemEventPayload struct {
	Coordinate coordinateviewmodel.ViewModel `json:"coordinate"`
	ItemId     string                        `json:"itemId"`
	ActionedAt time.Time                     `json:"actionedAt"`
}
type BuildItemEvent struct {
	Type    EventType             `json:"type"`
	Payload BuildItemEventPayload `json:"payload"`
}

type DestroyItemEventPayload struct {
	Coordinate coordinateviewmodel.ViewModel `json:"coordinate"`
	ActionedAt time.Time                     `json:"actionedAt"`
}
type DestroyItemEvent struct {
	Type    EventType               `json:"type"`
	Payload DestroyItemEventPayload `json:"payload"`
}

type ZoomAreaRequestedEventPayload struct {
	Area       areaviewmodel.ViewModel `json:"area"`
	ActionedAt time.Time               `json:"actionedAt"`
}
type ZoomAreaRequestedEvent struct {
	Type    EventType                     `json:"type"`
	Payload ZoomAreaRequestedEventPayload `json:"payload"`
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

func (presenter *Presenter) PresentInformationUpdatedEvent(dimension dimensionviewmodel.ViewModel) InformationUpdatedEvent {
	return InformationUpdatedEvent{
		Type: InformationUpdatedEventType,
		Payload: InformationUpdatedEventPayload{
			Dimension: dimension,
		},
	}
}

func (presenter *Presenter) PresentZoomedAreaUpdatedEvent(area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel) ZoomedAreaUpdatedEvent {
	return ZoomedAreaUpdatedEvent{
		Type: ZoomedAreaUpdatedEventType,
		Payload: ZoomedAreaUpdatedEventPayload{
			Area:      area,
			UnitBlock: unitBlock,
			UpdatedAt: time.Now(),
		},
	}
}

func (presenter *Presenter) PresentAreaZoomedEvent(area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel) AreaZoomedEvent {
	return AreaZoomedEvent{
		Type: AreaZoomedEventType,
		Payload: AreaZoomedEventPayload{
			Area:      area,
			UnitBlock: unitBlock,
		},
	}
}

func (presenter *Presenter) ParseBuildItemEvent(msg []byte) (commonmodel.Coordinate, itemmodel.ItemId, error) {
	var action BuildItemEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return commonmodel.Coordinate{}, itemmodel.ItemId{}, err
	}
	coordinate, err := action.Payload.Coordinate.ToValueObject()
	if err != nil {
		return commonmodel.Coordinate{}, itemmodel.ItemId{}, err
	}
	itemId, err := itemmodel.NewItemId(action.Payload.ItemId)
	if err != nil {
		return commonmodel.Coordinate{}, itemmodel.ItemId{}, err
	}

	return coordinate, itemId, nil
}

func (presenter *Presenter) ParseDestroyItemEvent(msg []byte) (commonmodel.Coordinate, error) {
	var action BuildItemEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return commonmodel.Coordinate{}, err
	}
	coordinate, err := action.Payload.Coordinate.ToValueObject()
	if err != nil {
		return commonmodel.Coordinate{}, err
	}

	return coordinate, nil
}

func (presenter *Presenter) ParseZoomAreaEvent(msg []byte) (commonmodel.Area, error) {
	var action ZoomAreaRequestedEvent
	err := json.Unmarshal(msg, &action)
	if err != nil {
		return commonmodel.Area{}, err
	}
	area, err := action.Payload.Area.ToValueObject()
	if err != nil {
		return commonmodel.Area{}, err
	}

	return area, nil
}
