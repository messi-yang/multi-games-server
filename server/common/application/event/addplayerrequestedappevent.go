package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
)

type AddPlayerRequestedAppEvent struct {
	LiveGameId string `json:"liveGameId"`
	PlayerId   string `json:"playerId"`
}

func NewAddPlayerRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) *AddPlayerRequestedAppEvent {
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

func (event *AddPlayerRequestedAppEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := gamecommonmodel.NewPlayerId(event.PlayerId)
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}
