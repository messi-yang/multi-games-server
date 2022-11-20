package redissubscriber

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/notification/commonredisnotification"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type RedisAddPlayerRequestedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
}

func NewRedisAddPlayerRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) RedisAddPlayerRequestedIntegrationEvent {
	return RedisAddPlayerRequestedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
	}
}

var RedisAddPlayerRequestedSubscriberChannel string = "add-player-requested"

type RedisAddPlayerRequestedSubscriber struct {
	redisProvider *commonredisnotification.RedisProvider
}

func NewRedisAddPlayerRequestedSubscriber() (notification.NotificationSubscriber[RedisAddPlayerRequestedIntegrationEvent], error) {
	return &RedisAddPlayerRequestedSubscriber{
		redisProvider: commonredisnotification.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisAddPlayerRequestedSubscriber) Subscribe(handler func(RedisAddPlayerRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisAddPlayerRequestedSubscriberChannel, func(message []byte) {
		var event RedisAddPlayerRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
