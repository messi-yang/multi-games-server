package dto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonjsondto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
)

type RedisZoomAreaRequestedEvent struct {
	LiveGameId commonjsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   commonjsondto.PlayerIdJsonDto   `json:"playerId"`
	Area       commonjsondto.AreaJsonDto       `json:"area"`
}

func NewRedisZoomAreaRequestedEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) RedisZoomAreaRequestedEvent {
	return RedisZoomAreaRequestedEvent{
		LiveGameId: commonjsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   commonjsondto.NewPlayerIdJsonDto(playerId),
		Area:       commonjsondto.NewAreaJsonDto(area),
	}
}

func NewRedisZoomAreaRequestedEventChannel() string {
	return "zoom-area-requested"
}

func (event *RedisZoomAreaRequestedEvent) GetLiveGameId() (livegamemodel.LiveGameId, error) {
	liveGameId, err := event.LiveGameId.ToValueObject()
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return liveGameId, nil
}

func (event *RedisZoomAreaRequestedEvent) GetPlayerId() (gamecommonmodel.PlayerId, error) {
	playerId, err := event.PlayerId.ToValueObject()
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return playerId, nil
}

func (event *RedisZoomAreaRequestedEvent) GetArea() (gamecommonmodel.Area, error) {
	area, err := event.Area.ToValueObject()
	if err != nil {
		return gamecommonmodel.Area{}, err
	}
	return area, nil
}
