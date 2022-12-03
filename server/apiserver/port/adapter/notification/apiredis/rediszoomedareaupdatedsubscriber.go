package apiredis

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/apiserver/port/adapter/presenter/presenterdto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
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

func RedisZoomedAreaUpdatedSubscriberChannel(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) string {
	return fmt.Sprintf("area-zoomed-live-game-id-%s-player-id-%s", liveGameId.GetId().String(), playerId.GetId().String())
}

type RedisZoomedAreaUpdatedSubscriber struct {
	liveGameId    livegamemodel.LiveGameId
	playerId      gamecommonmodel.PlayerId
	redisProvider *commonredis.RedisProvider
}

func NewRedisZoomedAreaUpdatedSubscriber(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) (commonnotification.NotificationSubscriber[RedisZoomedAreaUpdatedIntegrationEvent], error) {
	return &RedisZoomedAreaUpdatedSubscriber{
		liveGameId:    liveGameId,
		playerId:      playerId,
		redisProvider: commonredis.NewRedisProvider(),
	}, nil
}

func (subscriber *RedisZoomedAreaUpdatedSubscriber) Subscribe(handler func(RedisZoomedAreaUpdatedIntegrationEvent)) func() {
	unsubscriber := subscriber.redisProvider.Subscribe(RedisZoomedAreaUpdatedSubscriberChannel(subscriber.liveGameId, subscriber.playerId), func(message []byte) {
		var redisZoomedAreaUpdatedIntegrationEvent RedisZoomedAreaUpdatedIntegrationEvent
		json.Unmarshal(message, &redisZoomedAreaUpdatedIntegrationEvent)
		handler(redisZoomedAreaUpdatedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
