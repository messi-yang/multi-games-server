package dto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type RedisAddPlayerRequestedEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
}

func NewRedisAddPlayerRequestedEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) RedisAddPlayerRequestedEvent {
	return RedisAddPlayerRequestedEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
	}
}

func NewRedisAddPlayerRequestedEventChannel() string {
	return "add-player-requested"
}

func (event *RedisAddPlayerRequestedEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := event.LiveGameId.ToValueObject()
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *RedisAddPlayerRequestedEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := event.PlayerId.ToValueObject()
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}
