package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
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

var RedisReviveUnitsRequestedListenerChannel string = "revive-units-requested"

type RedisReviveUnitsRequestedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisReviveUnitsRequestedListenerConfiguration func(listener *RedisReviveUnitsRequestedListener) error

func NewRedisReviveUnitsRequestedListener(cfgs ...redisRedisReviveUnitsRequestedListenerConfiguration) (*RedisReviveUnitsRequestedListener, error) {
	t := &RedisReviveUnitsRequestedListener{
		redisInfrastructureService: service.NewRedisInfrastructureService(),
	}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (listener *RedisReviveUnitsRequestedListener) Subscribe(subscriber func(RedisReviveUnitsRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisReviveUnitsRequestedListenerChannel, func(message []byte) {
		var event RedisReviveUnitsRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
