package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
)

type AddPlayerRequestedAppEvent struct {
	LiveGameId string `json:"liveGameId"`
	PlayerId   string `json:"playerId"`
}

func NewAddPlayerRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) *AddPlayerRequestedAppEvent {
	return &AddPlayerRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
	}
}

func DeserializeAddPlayerRequestedAppEvent(message []byte) AddPlayerRequestedAppEvent {
	var event AddPlayerRequestedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewAddPlayerRequestedAppEventChannel() string {
	return "add-player-requested"
}

func (event *AddPlayerRequestedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *AddPlayerRequestedAppEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *AddPlayerRequestedAppEvent) GetPlayerId() (commonmodel.PlayerId, error) {
	playerId, err := commonmodel.NewPlayerId(event.PlayerId)
	if err != nil {
		return commonmodel.PlayerId{}, err
	}
	return playerId, nil
}
