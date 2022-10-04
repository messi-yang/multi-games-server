package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type zoomAreaRequestedIntegrationEventPayload struct {
	PlayerId uuid.UUID   `json:"playerId"`
	Area     dto.AreaDto `json:"area"`
}

type ZoomAreaRequestedIntegrationEvent struct {
	Payload   zoomAreaRequestedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                `json:"timestamp"`
}

func NewZoomAreaRequestedIntegrationEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-zoom-area-requested", gameId)
}

func NewZoomAreaRequestedIntegrationEvent(playerId uuid.UUID, area valueobject.Area) []byte {
	areaDto := dto.NewAreaDto(area)

	event := ZoomAreaRequestedIntegrationEvent{
		Payload: zoomAreaRequestedIntegrationEventPayload{
			PlayerId: playerId,
			Area:     areaDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
