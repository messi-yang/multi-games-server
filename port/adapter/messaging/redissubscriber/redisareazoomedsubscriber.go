package redissubscriber

import (
	"encoding/json"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/port/adapter/messaging/commonredissubscriber"
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

func RedisAreaZoomedSubscriberChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisAreaZoomedSubscriber struct {
	liveGameId             livegamemodel.LiveGameId
	playerId               gamecommonmodel.PlayerId
	redisMessageSubscriber *commonredissubscriber.RedisMessageSubscriber
}

func NewRedisAreaZoomedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonredissubscriber.RedisSubscriber[RedisAreaZoomedIntegrationEvent], error) {
	return &RedisAreaZoomedSubscriber{
		liveGameId:             liveGameId,
		playerId:               playerId,
		redisMessageSubscriber: commonredissubscriber.NewRedisMessageSubscriber(),
	}, nil
}

func (subscriber *RedisAreaZoomedSubscriber) Subscribe(handler func(RedisAreaZoomedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisMessageSubscriber.Subscribe(RedisAreaZoomedSubscriberChannel(subscriber.liveGameId, subscriber.playerId), func(message []byte) {
		var event RedisAreaZoomedIntegrationEvent
		json.Unmarshal(message, &event)
		handler(event)
	})

	return func() {
		unsubscriber()
	}
}
