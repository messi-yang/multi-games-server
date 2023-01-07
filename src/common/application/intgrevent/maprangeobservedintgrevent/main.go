package maprangeobservedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/maprangeviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitmapviewmodel"
)

type Event struct {
	Name       string                      `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	PlayerId   string                      `json:"playerId"`
	MapRange   maprangeviewmodel.ViewModel `json:"mapRange"`
	UnitMap    unitmapviewmodel.ViewModel  `json:"unitMap"`
}

var EVENT_NAME = "MAP_RANGE_OBSERVED"

func New(liveGameId string, playerId string, mapRange maprangeviewmodel.ViewModel, unitMap unitmapviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapRange:   mapRange,
		UnitMap:    unitMap,
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
