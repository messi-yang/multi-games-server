package gameredis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/presenter/presenterdto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisZoomAreaRequestedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
	Area       presenterdto.AreaPresenterDto       `json:"area"`
}

func NewRedisZoomAreaRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) RedisZoomAreaRequestedIntegrationEvent {
	return RedisZoomAreaRequestedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
		Area:       presenterdto.NewAreaPresenterDto(area),
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
