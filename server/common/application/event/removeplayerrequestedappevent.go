package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type RemovePlayerRequestedAppEvent struct {
	LiveGameId string                        `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto `json:"playerId"`
}

func NewRemovePlayerRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) *RemovePlayerRequestedAppEvent {
	return &RemovePlayerRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
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

func (event *RemovePlayerRequestedAppEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := event.PlayerId.ToValueObject()
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}
