package removeplayerrequestedintegrationevent

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type Event struct {
	Name       string `json:"name"`
	LiveGameId string `json:"liveGameId"`
	PlayerId   string `json:"playerId"`
}

var EVENT_NAME = "REMOVE_PLAYER_REQUESTED"

func New(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) *Event {
	return &Event{
		Name:       EVENT_NAME,
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
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

func (event *Event) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *Event) GetPlayerId() (commonmodel.PlayerId, error) {
	playerId, err := commonmodel.NewPlayerId(event.PlayerId)
	if err != nil {
		return commonmodel.PlayerId{}, err
	}
	return playerId, nil
}
