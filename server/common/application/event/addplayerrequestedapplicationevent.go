package event

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type AddPlayerRequestedApplicationEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
}

func NewAddPlayerRequestedApplicationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) ApplicationEvent {
	return &AddPlayerRequestedApplicationEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
	}
}

func NewAddPlayerRequestedApplicationEventChannel() string {
	return "add-player-requested"
}

func (event *AddPlayerRequestedApplicationEvent) Serialize() []byte {
	message, _ := json.Marshal(event)
	return message
}

func (event *AddPlayerRequestedApplicationEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := event.LiveGameId.ToValueObject()
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *AddPlayerRequestedApplicationEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := event.PlayerId.ToValueObject()
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}
