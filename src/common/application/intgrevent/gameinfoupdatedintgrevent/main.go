package gameinfoupdatedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/dimensionviewmodel"
)

type Event struct {
	Name       string                       `json:"name"`
	LiveGameId string                       `json:"liveGameId"`
	PlayerId   string                       `json:"playerId"`
	Dimension  dimensionviewmodel.ViewModel `json:"dimension"`
}

var EVENT_NAME = "GAME_INFO_UPDATED"

func New(liveGameId string, playerId string, dimension dimensionviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Dimension:  dimension,
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
