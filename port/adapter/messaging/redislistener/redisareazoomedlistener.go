package redislistener

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredislistener"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamecommonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
)

type RedisAreaZoomedIntegrationEvent struct {
	LiveGameId presenterdto.LiveGameIdPresenterDto `json:"liveGameId"`
	PlayerId   presenterdto.PlayerIdPresenterDto   `json:"playerId"`
	Area       presenterdto.AreaPresenterDto       `json:"area"`
	UnitBlock  presenterdto.UnitBlockPresenterDto  `json:"unitBlock"`
}

func NewRedisAreaZoomedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) RedisAreaZoomedIntegrationEvent {
	return RedisAreaZoomedIntegrationEvent{
		LiveGameId: presenterdto.NewLiveGameIdPresenterDto(liveGameId),
		PlayerId:   presenterdto.NewPlayerIdPresenterDto(playerId),
		Area:       presenterdto.NewAreaPresenterDto(area),
		UnitBlock:  presenterdto.NewUnitBlockPresenterDto(unitBlock),
	}
}

func RedisAreaZoomedListenerChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisAreaZoomedListener struct {
	liveGameId             livegamemodel.LiveGameId
	playerId               gamecommonmodel.PlayerId
	redisMessageSubscriber *commonredislistener.RedisMessageSubscriber
}

func NewRedisAreaZoomedListener(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonredislistener.RedisListener[RedisAreaZoomedIntegrationEvent], error) {
	return &RedisAreaZoomedListener{
		liveGameId:             liveGameId,
		playerId:               playerId,
		redisMessageSubscriber: commonredislistener.NewRedisMessageSubscriber(),
	}, nil
}

func (listener *RedisAreaZoomedListener) Subscribe(subscriber func(RedisAreaZoomedIntegrationEvent)) func() {
	unsubscriber := listener.redisMessageSubscriber.Subscribe(RedisAreaZoomedListenerChannel(listener.liveGameId, listener.playerId), func(message []byte) {
		var event RedisAreaZoomedIntegrationEvent
		json.Unmarshal(message, &event)
		subscriber(event)
	})

	return func() {
		unsubscriber()
	}
}
