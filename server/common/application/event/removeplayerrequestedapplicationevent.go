package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type RemovePlayerRequestedApplicationEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
}

func NewRemovePlayerRequestedApplicationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) *RemovePlayerRequestedApplicationEvent {
	return &RemovePlayerRequestedApplicationEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
	}
}

func DeserializeRemovePlayerRequestedApplicationEvent(message []byte) RemovePlayerRequestedApplicationEvent {
	var event RemovePlayerRequestedApplicationEvent
	json.Unmarshal(message, &event)
	return event
}

func NewRemovePlayerRequestedApplicationEventChannel() string {
	return "remove-player-requested"
}

func (event *RemovePlayerRequestedApplicationEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *RemovePlayerRequestedApplicationEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := event.LiveGameId.ToValueObject()
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *RemovePlayerRequestedApplicationEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := event.PlayerId.ToValueObject()
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}
