package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/adapter/common/dto/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type DestroyItemRequestedAppEvent struct {
	LiveGameId string                    `json:"liveGameId"`
	Coordinate jsondto.CoordinateJsonDto `json:"coordinate"`
}

func NewDestroyItemRequestedAppEvent(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) *DestroyItemRequestedAppEvent {
	return &DestroyItemRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		Coordinate: jsondto.NewCoordinateJsonDto(coordinate),
	}
}

func DeserializeDestroyItemRequestedAppEvent(message []byte) DestroyItemRequestedAppEvent {
	var event DestroyItemRequestedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewDestroyItemRequestedAppEventChannel() string {
	return "destroy-item-requested"
}

func (event *DestroyItemRequestedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *DestroyItemRequestedAppEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *DestroyItemRequestedAppEvent) GetCoordinate() (commonmodel.Coordinate, error) {
	coordinate, err := event.Coordinate.ToValueObject()
	if err != nil {
		return commonmodel.Coordinate{}, err
	}
	return coordinate, nil
}
