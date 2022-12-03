package gameredis

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisZoomAreaRequestedIntegrationEvent struct {
	LiveGameId jsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   jsondto.PlayerIdJsonDto   `json:"playerId"`
	Area       jsondto.AreaJsonDto       `json:"area"`
}

func NewRedisZoomAreaRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) RedisZoomAreaRequestedIntegrationEvent {
	return RedisZoomAreaRequestedIntegrationEvent{
		LiveGameId: jsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   jsondto.NewPlayerIdJsonDto(playerId),
		Area:       jsondto.NewAreaJsonDto(area),
	}
}

var RedisZoomAreaRequestedSubscriberChannel string = "zoom-area-requested"

type RedisZoomAreaRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisZoomAreaRequestedSubscriber() (commonnotification.NotificationSubscriber[RedisZoomAreaRequestedIntegrationEvent], error) {
	return &RedisZoomAreaRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisZoomAreaRequestedSubscriber) Subscribe(handler func(RedisZoomAreaRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisZoomAreaRequestedSubscriberChannel, func(message []byte) {
		var event RedisZoomAreaRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
