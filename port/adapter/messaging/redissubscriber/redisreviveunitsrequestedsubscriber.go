package redissubscriber

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredissubscriber"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type RedisReviveUnitsRequestedIntegrationEvent struct {
	LiveGameId  presenterdto.LiveGameIdPresenterDto   `json:"liveGameId"`
	Coordinates []presenterdto.CoordinatePresenterDto `json:"coordinates"`
}

func NewRedisReviveUnitsRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) RedisReviveUnitsRequestedIntegrationEvent {
	return RedisReviveUnitsRequestedIntegrationEvent{
		LiveGameId:  presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		Coordinates: presenterdto.NewCoordinatePresenterDtos(coordinates),
	}
}

var RedisReviveUnitsRequestedSubscriberChannel string = "revive-units-requested"

type RedisReviveUnitsRequestedSubscriber struct {
	redisMessageSubscriber *commonredissubscriber.RedisMessageSubscriber
}

func NewRedisReviveUnitsRequestedSubscriber() (commonredissubscriber.RedisSubscriber[RedisReviveUnitsRequestedIntegrationEvent], error) {
	return &RedisReviveUnitsRequestedSubscriber{
		redisMessageSubscriber: commonredissubscriber.NewRedisMessageSubscriber(),
	}, nil
}

func (subscriber *RedisReviveUnitsRequestedSubscriber) Subscribe(handler func(RedisReviveUnitsRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisMessageSubscriber.Subscribe(RedisReviveUnitsRequestedSubscriberChannel, func(message []byte) {
		var event RedisReviveUnitsRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
