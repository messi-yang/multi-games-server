package addplayerrequestedevent

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/dto/playerdto"
	"github.com/google/uuid"
)

type payload struct {
	Player playerdto.Dto `json:"player"`
}

type Event struct {
	Payload   payload   `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEventTopic(gameId uuid.UUID) string {
	return fmt.Sprintf("game-room-%s-add-player-requested", gameId)
}

func NewEvent(player entity.Player) []byte {
	playerDto := playerdto.ToDto(player)

	event := Event{
		Payload: payload{
			Player: playerDto,
		},
		Timestamp: time.Now(),
	}

	jsonBytes, _ := json.Marshal(event)

	return jsonBytes
}
