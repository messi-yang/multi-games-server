package gameinfoupdatedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/mapsizeviewmodel"
)

type Event struct {
	Name       string                     `json:"name"`
	LiveGameId string                     `json:"liveGameId"`
	PlayerId   string                     `json:"playerId"`
	MapSize    mapsizeviewmodel.ViewModel `json:"mapSize"`
}

var EVENT_NAME = "GAME_INFO_UPDATED"

func New(liveGameId string, playerId string, mapSize mapsizeviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapSize:    mapSize,
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
