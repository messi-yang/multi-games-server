package apiredis

import (
	"encoding/json"
	"fmt"

	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/common/dto/jsondto"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
)

type RedisZoomedAreaUpdatedIntegrationEvent struct {
	LiveGameId jsondto.LiveGameIdJsonDto `json:"liveGameId"`
	PlayerId   jsondto.PlayerIdJsonDto   `json:"playerId"`
	Area       jsondto.AreaJsonDto       `json:"area"`
	UnitBlock  jsondto.UnitBlockJsonDto  `json:"unitBlock"`
}

func NewRedisZoomedAreaUpdatedIntegrationEvent(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area, unitBlock gamecommonmodel.UnitBlock) RedisZoomedAreaUpdatedIntegrationEvent {
	return RedisZoomedAreaUpdatedIntegrationEvent{
		LiveGameId: jsondto.NewLiveGameIdJsonDto(liveGameId),
		PlayerId:   jsondto.NewPlayerIdJsonDto(playerId),
		Area:       jsondto.NewAreaJsonDto(area),
		UnitBlock:  jsondto.NewUnitBlockJsonDto(unitBlock),
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
