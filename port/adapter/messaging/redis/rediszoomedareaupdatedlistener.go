package redis

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type RedisZoomedAreaUpdatedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
	Area       presenterdto.AreaPresenterDto       `json:"area"`
	UnitBlock  presenterdto.UnitBlockPresenterDto  `json:"unitBlock"`
}

func NewRedisZoomedAreaUpdatedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area presenterdto.AreaPresenterDto, unitBlock presenterdto.UnitBlockPresenterDto) RedisZoomedAreaUpdatedIntegrationEvent {
	return RedisZoomedAreaUpdatedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
		Area:       area,
		UnitBlock:  unitBlock,
	}
}

func RedisZoomedAreaUpdatedListenerChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisZoomedAreaUpdatedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisRedisZoomedAreaUpdatedListenerConfiguration func(listener *RedisZoomedAreaUpdatedListener) error

func NewRedisZoomedAreaUpdatedListener(cfgs ...redisRedisZoomedAreaUpdatedListenerConfiguration) (*RedisZoomedAreaUpdatedListener, error) {
	t := &RedisZoomedAreaUpdatedListener{
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

func (listener *RedisZoomedAreaUpdatedListener) Subscribe(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, subscriber func(RedisZoomedAreaUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(RedisZoomedAreaUpdatedListenerChannel(liveGameId, playerId), func(message []byte) {
		var redisZoomedAreaUpdatedIntegrationEvent RedisZoomedAreaUpdatedIntegrationEvent
		json.Unmarshal(message, &redisZoomedAreaUpdatedIntegrationEvent)
		subscriber(redisZoomedAreaUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
