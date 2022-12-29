package gameinfoupdatedintegrationevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/interface/viewmodel/dimensionviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type Event struct {
	Name       string                                `json:"name"`
	LiveGameId string                                `json:"liveGameId"`
	PlayerId   string                                `json:"playerId"`
	Dimension  dimensionviewmodel.DimensionViewModel `json:"dimension"`
}

var EVENT_NAME = "GAME_INFO_UPDATED"

func New(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, dimension commonmodel.Dimension) *Event {
	return &Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
		Dimension:  dimensionviewmodel.New(dimension),
	}
}

func Deserialize(message []byte) Event {
	var event Event
	json.Unmarshal(message, &event)
	return event
}

func (event *Event) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *Event) GetDimension() (commonmodel.Dimension, error) {
	dimension, err := event.Dimension.ToValueObject()
	if err != nil {
		return commonmodel.Dimension{}, nil
	}
	return dimension, nil
}
