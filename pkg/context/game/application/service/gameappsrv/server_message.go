package gameappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/dto"
	"github.com/google/uuid"
)

type UnitCreatedServerMessage struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type UnitDeletedServerMessage struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type PlayerJoinedServerMessage struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type PlayerLeftServerMessage struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type PlayerMovedServerMessage struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

func NewWorldServerMessageChannel(worldIdDto uuid.UUID) string {
	return fmt.Sprintf("WORLD_%s_CHANNEL", worldIdDto)
}
