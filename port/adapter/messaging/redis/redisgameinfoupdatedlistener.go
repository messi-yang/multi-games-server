package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type RedisGameInfoUpdatedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
	Dimension  presenterdto.DimensionPresenterDto  `json:"dimension"`
}

func NewRedisGameInfoUpdatedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, dimension gamecommonmodel.Dimension) RedisGameInfoUpdatedIntegrationEvent {
	return RedisGameInfoUpdatedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
		Dimension:  presenterdto.NewDimensionPresenterDto(dimension),
	}
}

func RedisGameInfoUpdatedListenerChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("game-info-updated-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisGameInfoUpdatedListener struct {
	redisMessageSubscriber *commonredis.RedisMessageSubscriber
}

type redisRedisGameInfoUpdatedListenerConfiguration func(listener *RedisGameInfoUpdatedListener) error

func NewRedisGameInfoUpdatedListener(cfgs ...redisRedisGameInfoUpdatedListenerConfiguration) (*RedisGameInfoUpdatedListener, error) {
	t := &RedisGameInfoUpdatedListener{
		redisMessageSubscriber: commonredis.NewRedisMessageSubscriber(),
	}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (listener *RedisGameInfoUpdatedListener) Subscribe(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, subscriber func(RedisGameInfoUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisMessageSubscriber.Subscribe(RedisGameInfoUpdatedListenerChannel(liveGameId, playerId), func(message []byte) {
		var event RedisGameInfoUpdatedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
