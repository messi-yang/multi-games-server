package zoommaprangerequestedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/maprangeviewmodel"
)

type Event struct {
	Name       string                      `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	PlayerId   string                      `json:"playerId"`
	MapRange   maprangeviewmodel.ViewModel `json:"mapRange"`
}

var EVENT_NAME = "ZOOM_MAP_RANGE_REQUESTED"

func New(liveGameId string, playerId string, mapRange maprangeviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapRange:   mapRange,
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
