package event

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/livegamemodel"
)

type RemovePlayerRequestedAppEvent struct {
	LiveGameId string `json:"liveGameId"`
	PlayerId   string `json:"playerId"`
}

func NewRemovePlayerRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) *RemovePlayerRequestedAppEvent {
	return &RemovePlayerRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   playerId.ToString(),
	}
}

func DeserializeRemovePlayerRequestedAppEvent(message []byte) RemovePlayerRequestedAppEvent {
	var event RemovePlayerRequestedAppEvent
	json.Unmarshal(message, &event)
	return event
}

func NewRemovePlayerRequestedAppEventChannel() string {
	return "remove-player-requested"
}

func (event *RemovePlayerRequestedAppEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *RemovePlayerRequestedAppEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := livegamemodel.NewLiveGameId(event.LiveGameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *RemovePlayerRequestedAppEvent) GetPlayerId() (commonmodel.PlayerId, error) {
	playerId, err := commonmodel.NewPlayerId(event.PlayerId)
	if err != nil {
		return commonmodel.PlayerId{}, err
	}
	return playerId, nil
}
