package observedmaprangeupdatedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/gamemapviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/maprangeviewmodel"
)

type Event struct {
	Name       string                      `json:"name"`
	LiveGameId string                      `json:"liveGameId"`
	PlayerId   string                      `json:"playerId"`
	MapRange   maprangeviewmodel.ViewModel `json:"mapRange"`
	GameMap    gamemapviewmodel.ViewModel  `json:"gameMap"`
}

var EVENT_NAME = "OBSERVED_MAP_RANGE_UPDATED"

func New(liveGameId string, playerId string, mapRange maprangeviewmodel.ViewModel, gameMap gamemapviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		MapRange:   mapRange,
		GameMap:    gameMap,
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
