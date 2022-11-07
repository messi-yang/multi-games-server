package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/dto"
	"github.com/google/uuid"
)

type zoomAreaRequestedIntegrationEventPayload struct {
	PlayerId dto.PlayerIdDto `json:"playerId"`
	Area     dto.AreaDto     `json:"area"`
}

type ZoomAreaRequestedIntegrationEvent struct {
	Payload   zoomAreaRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                `json:"timestamp"`
}

func NewZoomAreaRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-zoom-area-requested", gameId)
}

func NewZoomAreaRequestedIntegrationEvent(playerIdDto dto.PlayerIdDto, areaDto dto.AreaDto) []byte {
	event := ZoomAreaRequestedIntegrationEvent{
		Payload: zoomAreaRequestedIntegrationEventPayload{
			PlayerId: playerIdDto,
			Area:     areaDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
