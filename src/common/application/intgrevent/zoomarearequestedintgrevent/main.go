package zoomarearequestedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
)

type Event struct {
	Name       string                  `json:"name"`
	LiveGameId string                  `json:"liveGameId"`
	PlayerId   string                  `json:"playerId"`
	Area       areaviewmodel.ViewModel `json:"area"`
}

var EVENT_NAME = "ZOOM_AREA_REQUESTED"

func New(liveGameId string, playerId string, area areaviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Area:       area,
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
