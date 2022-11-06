package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/google/uuid"
)

type zoomedAreaUpdatedIntegrationEventPayload struct {
	Area      dto.AreaDto      `json:"area"`
	UnitBlock dto.UnitBlockDto `json:"unitBlock"`
}

type ZoomedAreaUpdatedIntegrationEvent struct {
	Payload   zoomedAreaUpdatedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                                `json:"timestamp"`
}

func NewZoomedAreaUpdatedIntegrationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-zoomed-area-updated", gameId, playerId)
}

func NewZoomedAreaUpdatedIntegrationEvent(area valueobject.Area, unitBlock valueobject.UnitBlock) []byte {
	areaDto := dto.NewAreaDto(area)
	unitBlockDto := dto.Dto(unitBlock)

	event := ZoomedAreaUpdatedIntegrationEvent{
		Payload: zoomedAreaUpdatedIntegrationEventPayload{
			Area:      areaDto,
			UnitBlock: unitBlockDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
