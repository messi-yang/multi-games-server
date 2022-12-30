package addplayerrequestedintgrevent

import (
	"encoding/json"
)

type Event struct {
	Name       string `json:"name"`
	LiveGameId string `json:"liveGameId"`
	PlayerId   string `json:"playerId"`
}

var EVENT_NAME = "ADD_PLAYER_REQUESTED"

func New(liveGameId string, playerId string) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
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
