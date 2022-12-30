package destroyitemrequestedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
)

type Event struct {
	Name       string                        `json:"name"`
	LiveGameId string                        `json:"liveGameId"`
	Coordinate coordinateviewmodel.ViewModel `json:"coordinate"`
}

var EVENT_NAME = "DESTROY_ITEM_REQUESTED"

func New(liveGameId string, coordinate coordinateviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		Coordinate: coordinate,
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
