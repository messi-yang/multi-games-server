package areazoomedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type payload struct {
	Area    dto.AreaDto    `json:"area"`
	UnitMap dto.UnitMapDto `json:"unitMap"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID, playerId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-player-%s-area-zoomed", gameId, playerId)
}

func NewEvent(area valueobject.Area, unitMap valueobject.UnitMap) []byte {
	areaDto := dto.NewAreaDto(area)
	unitMapDto := dto.NewUnitMapDto(&unitMap)

	event := Event{
		Payload: payload{
			Area:    areaDto,
			UnitMap: unitMapDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
