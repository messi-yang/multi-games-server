package buliditemrequestedintegrationevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type Event struct {
	Name       string                                  `json:"name"`
	LiveGameId string                                  `json:"liveGameId"`
	Coordinate coordinateviewmodel.CoordinateViewModel `json:"coordinate"`
	ItemId     string                                  `json:"coordinates"`
}

var EVENT_NAME = "BUILD_ITEM_REQUESTED"

func New(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) *Event {
	return &Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId.ToString(),
		Coordinate: coordinateviewmodel.New(coordinate),
		ItemId:     itemId.ToString(),
	}
}

func Deserialize(message []byte) Event {
	var event Event
	json.Unmarshal(message, &event)
	return event
}

func (event *Event) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *Event) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *Event) GetItemId() (itemmodel.ItemId, error) {
	itemId, err := itemmodel.NewItemId(event.ItemId)
	if err != nil {
		return itemmodel.ItemId{}, err
	}
	return itemId, nil
}

func (event *Event) GetCoordinate() (commonmodel.Coordinate, error) {
	coordinate, err := event.Coordinate.ToValueObject()
	if err != nil {
		return commonmodel.Coordinate{}, err
	}
	return coordinate, nil
}
