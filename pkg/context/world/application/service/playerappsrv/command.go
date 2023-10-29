package playerappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type ChangePlayerHeldItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	ItemId   uuid.UUID
}

type EnterWorldCommand struct {
	WorldId          uuid.UUID
	PlayerName       string
	PlayerHeldItemId uuid.UUID
}

type MakePlayerWalkCommand struct {
	WorldId        uuid.UUID
	PlayerId       uuid.UUID
	ActionPosition dto.PositionDto
	Direction      int8
}

type MakePlayerStandCommand struct {
	WorldId        uuid.UUID
	PlayerId       uuid.UUID
	ActionPosition dto.PositionDto
	Direction      int8
}

type SendPlayerIntoPortalCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	Position dto.PositionDto
}

type LeaveWorldCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
