package gameredis

import (
	"encoding/json"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisReviveUnitsRequestedIntegrationEvent struct {
	LiveGameId  jsondto.LiveGameIdJsonDto   `json:"liveGameId"`
	Coordinates []jsondto.CoordinateJsonDto `json:"coordinates"`
}

func NewRedisReviveUnitsRequestedIntegrationEvent(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) RedisReviveUnitsRequestedIntegrationEvent {
	return RedisReviveUnitsRequestedIntegrationEvent{
		LiveGameId:  jsondto.NewLiveGameIdJsonDto(liveGameId),
		Coordinates: jsondto.NewCoordinateJsonDtos(coordinates),
	}
}

var RedisReviveUnitsRequestedSubscriberChannel string = "revive-units-requested"

type RedisReviveUnitsRequestedSubscriber struct {
	redisProvider *commonredis.RedisProvider
}

func NewRedisReviveUnitsRequestedSubscriber() (commonnotification.NotificationSubscriber[RedisReviveUnitsRequestedIntegrationEvent], error) {
	return &RedisReviveUnitsRequestedSubscriber{
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisReviveUnitsRequestedSubscriber) Subscribe(handler func(RedisReviveUnitsRequestedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisReviveUnitsRequestedSubscriberChannel, func(message []byte) {
		var event RedisReviveUnitsRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
