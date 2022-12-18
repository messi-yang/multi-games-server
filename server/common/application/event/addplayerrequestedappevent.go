package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type AddPlayerRequestedAppEvent struct {
	LiveGameId string                        `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto `json:"playerId"`
}

func NewAddPlayerRequestedAppEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) *AddPlayerRequestedAppEvent {
	return &AddPlayerRequestedAppEvent{
		LiveGameId: liveGameId.ToString(),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
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
	playerId, err := event.PlayerId.ToValueObject()
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}
