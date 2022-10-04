package addplayerrequestedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto"
	"github.com/google/uuid"
)

type payload struct {
	Player dto.PlayerDto `json:"player"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewEvent(player entity.Player) []byte {
	playerDto := dto.NewPlayerDto(player)

	event := Event{
		Payload: payload{
			Player: playerDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
