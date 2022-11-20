package redissubscriber

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/notification/commonredisnotification"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
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
	redisProvider *commonredisnotification.RedisProvider
}

func NewRedisZoomAreaRequestedSubscriber() (notification.NotificationSubscriber[RedisZoomAreaRequestedIntegrationEvent], error) {
	return &RedisZoomAreaRequestedSubscriber{
		redisProvider: commonredisnotification.NewRedisProvider(),
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
