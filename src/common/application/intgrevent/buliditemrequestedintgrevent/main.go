package buliditemrequestedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/locationviewmodel"
)

type Event struct {
	Name       string                      `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	Location   locationviewmodel.ViewModel `json:"location"`
	ItemId     string                      `json:"locations"`
}

var EVENT_NAME = "BUILD_ITEM_REQUESTED"

func New(liveGameId string, location locationviewmodel.ViewModel, itemId string) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		Location:   location,
		ItemId:     itemId,
	}
}

func Deserialize(message []byte) Event {
	var event Event
	json.Unmarshal(message, &event)
	return event
}

func (event Event) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}
