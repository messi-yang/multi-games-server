package integrationevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type areaZoomedIntegrationEventPayload struct {
	Area    dto.AreaDto    `json:"area"`
	UnitMap dto.UnitMapDto `json:"unitMap"`
}

type AreaZoomedIntegrationEvent struct {
	Payload   areaZoomedIntegrationEventPayload `json:"payload"`
	Timestamp time.Time                         `json:"timestamp"`
}

func NewAreaZoomedIntegrationEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", gameId, playerId)
}

func NewAreaZoomedIntegrationEvent(area valueobject.Area, unitMap valueobject.UnitMap) []byte {
	areaDto := dto.NewAreaDto(area)
	unitMapDto := dto.Dto(unitMap)

	event := AreaZoomedIntegrationEvent{
		Payload: areaZoomedIntegrationEventPayload{
			Area:    areaDto,
			UnitMap: unitMapDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
