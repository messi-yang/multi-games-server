package apiredis

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisAreaZoomedIntegrationEvent struct {
	LiveGameId jsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   jsondto.PlayerIdJsonDto   `json:"playerId"`
	Area       jsondto.AreaJsonDto       `json:"area"`
	UnitBlock  jsondto.UnitBlockJsonDto  `json:"unitBlock"`
}

func NewRedisAreaZoomedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) RedisAreaZoomedIntegrationEvent {
	return RedisAreaZoomedIntegrationEvent{
		LiveGameId: jsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   jsondto.NewPlayerIdJsonDto(playerId),
		Area:       jsondto.NewAreaJsonDto(area),
		UnitBlock:  jsondto.NewUnitBlockJsonDto(unitBlock),
	}
}

func RedisAreaZoomedSubscriberChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisAreaZoomedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      gamecommonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisAreaZoomedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonnotification.NotificationSubscriber[RedisAreaZoomedIntegrationEvent], error) {
	return &RedisAreaZoomedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAreaZoomedSubscriber) Subscribe(handler func(RedisAreaZoomedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisAreaZoomedSubscriberChannel(subscriber.liveGameId, subscriber.playerId), func(message []byte) {
		var event RedisAreaZoomedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
