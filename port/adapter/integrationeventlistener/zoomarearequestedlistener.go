package integrationeventlistener

import (
	"encoding/json"

	"github.com/dum-dum-genius/game-of-liberty-computer/common/infrastructure/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/port/adapter/presenter/presenterdto"
	"github.com/google/uuid"
)

type ZoomAreaRequestedIntegrationEvent struct {
	GameId   uuid.UUID                     `json:"gameId"`
	PlayerId uuid.UUID                     `json:"playerId"`
	Area     presenterdto.AreaPresenterDto `json:"area"`
}

func NewZoomAreaRequestedIntegrationEvent(gameId uuid.UUID, playerId uuid.UUID, areaPresenterDto presenterdto.AreaPresenterDto) ZoomAreaRequestedIntegrationEvent {
	return ZoomAreaRequestedIntegrationEvent{
		GameId:   gameId,
		PlayerId: playerId,
		Area:     areaPresenterDto,
	}
}

var ZoomAreaRequestedListenerChannel string = "zoom-area-requested"

type ZoomAreaRequestedListener struct {
	redisInfrastructureService *service.RedisInfrastructureService
}

type redisZoomedAreaRequestedListenerConfiguration func(listener *ZoomAreaRequestedListener) error

func NewZoomAreaRequestedListener(cfgs ...redisZoomedAreaRequestedListenerConfiguration) (*ZoomAreaRequestedListener, error) {
	t := &ZoomAreaRequestedListener{
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

func (listener *ZoomAreaRequestedListener) Subscribe(subscriber func(ZoomAreaRequestedIntegrationEvent)) func() {
	unsubscriber := listener.redisInfrastructureService.Subscribe(ZoomAreaRequestedListenerChannel, func(message []byte) {
		var zoomAreaRequestedIntegrationEvent ZoomAreaRequestedIntegrationEvent
		json.Unmarshal(message, &zoomAreaRequestedIntegrationEvent)
		subscriber(zoomAreaRequestedIntegrationEvent)
	})

	return func() {
		unsubscriber()
	}
}
