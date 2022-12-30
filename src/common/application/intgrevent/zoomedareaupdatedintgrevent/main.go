package zoomedareaupdatedintgrevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitblockviewmodel"
)

type Event struct {
	Name       string                       `json:"name"`
	LiveGameId string                       `json:"liveGameId"`
	PlayerId   string                       `json:"playerId"`
	Area       areaviewmodel.ViewModel      `json:"area"`
	UnitBlock  unitblockviewmodel.ViewModel `json:"unitBlock"`
}

var EVENT_NAME = "ZOOMED_AREA_UPDATED"

func New(liveGameId string, playerId string, area areaviewmodel.ViewModel, unitBlock unitblockviewmodel.ViewModel) Event {
	return Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId,
		PlayerId:   playerId,
		Area:       area,
		UnitBlock:  unitBlock,
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
