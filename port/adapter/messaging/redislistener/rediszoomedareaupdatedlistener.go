package redislistener

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredislistener"
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

func NewRedisZoomedAreaUpdatedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) RedisZoomedAreaUpdatedIntegrationEvent {
	return RedisZoomedAreaUpdatedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
		Area:       presenterdto.NewAreaPresenterDto(area),
		UnitBlock:  presenterdto.NewUnitBlockPresenterDto(unitBlock),
	}
}

func RedisZoomedAreaUpdatedListenerChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisZoomedAreaUpdatedListener struct {
	liveGameId             livegamemodel.LiveGameId
	playerId               gamecommonmodel.PlayerId
	redisMessageSubscriber *commonredislistener.RedisMessageSubscriber
}

type redisRedisZoomedAreaUpdatedListenerConfiguration func(listener *RedisZoomedAreaUpdatedListener) error

func NewRedisZoomedAreaUpdatedListener(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonredislistener.RedisListener[RedisZoomedAreaUpdatedIntegrationEvent], error) {
	return &RedisZoomedAreaUpdatedListener{
		liveGameId:             liveGameId,
		playerId:               playerId,
		redisMessageSubscriber: commonredislistener.NewRedisMessageSubscriber(),
	}, nil
}

func (listener *RedisZoomedAreaUpdatedListener) Subscribe(subscriber func(RedisZoomedAreaUpdatedIntegrationEvent)) func() {
	unsubscriber := listener.redisMessageSubscriber.Subscribe(RedisZoomedAreaUpdatedListenerChannel(listener.liveGameId, listener.playerId), func(message []byte) {
		var redisZoomedAreaUpdatedIntegrationEvent RedisZoomedAreaUpdatedIntegrationEvent
		json.Unmarshal(message, &redisZoomedAreaUpdatedIntegrationEvent)
		subscriber(redisZoomedAreaUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
