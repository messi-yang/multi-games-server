package redis

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredis"
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

var RedisZoomAreaRequestedListenerChannel string = "zoom-area-requested"

type RedisZoomAreaRequestedListener struct {
	redisMessageSubscriber *commonredis.RedisMessageSubscriber
}

type redisZoomedAreaRequestedListenerConfiguration func(listener *RedisZoomAreaRequestedListener) error

func NewRedisZoomAreaRequestedListener(cfgs ...redisZoomedAreaRequestedListenerConfiguration) (*RedisZoomAreaRequestedListener, error) {
	t := &RedisZoomAreaRequestedListener{
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

func (listener *RedisZoomAreaRequestedListener) Subscribe(subscriber func(RedisZoomAreaRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisMessageSubscriber.Subscribe(RedisZoomAreaRequestedListenerChannel, func(message []byte) {
		var event RedisZoomAreaRequestedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
