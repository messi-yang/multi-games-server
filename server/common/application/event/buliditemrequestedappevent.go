package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type BuildItemRequestedAppEvent struct {
	LiveGameId string                    `json:"liveGameId"`
	Coordinate jsondto.CoordinateJsonDto `json:"coordinate"`
	ItemId     string                    `json:"coordinates"`
}

func NewBuildItemRequestedAppEvent(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) *BuildItemRequestedAppEvent {
	return &BuildItemRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		Coordinate: jsondto.NewCoordinateJsonDto(coordinate),
		ItemId:     itemId.ToString(),
	}
}

func DeserializeBuildItemRequestedAppEvent(message []byte) BuildItemRequestedAppEvent {
	var event BuildItemRequestedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewBuildItemRequestedAppEventChannel() string {
	return "build-item-requested"
}

func (event *BuildItemRequestedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *BuildItemRequestedAppEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *BuildItemRequestedAppEvent) GetItemId() (itemmodel.ItemId, error) {
	itemId, err := itemmodel.NewItemId(event.ItemId)
	if err != nil {
		return itemmodel.ItemId{}, err
	}
	return itemId, nil
}

func (event *BuildItemRequestedAppEvent) GetCoordinate() (commonmodel.Coordinate, error) {
	coordinate, err := event.Coordinate.ToValueObject()
	if err != nil {
		return commonmodel.Coordinate{}, err
	}
	return coordinate, nil
}
