package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/dto"
	"github.com/google/uuid"
)

type areaZoomedIntegrationEventPayload struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
}

type AreaZoomedIntegrationEvent struct {
	Payload   areaZoomedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                         `json:"timestamp"`
}

func NewAreaZoomedIntegrationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", gameId, playerId)
}

func NewAreaZoomedIntegrationEvent(areaDto dto.AreaDto, unitBlockDto dto.UnitBlockDto) []byte {
	event := AreaZoomedIntegrationEvent{
		Payload: areaZoomedIntegrationEventPayload{
			Area:      areaDto,
			UnitBlock: unitBlockDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
