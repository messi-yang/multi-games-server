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

type ChangePlayerActionCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	Action   dto.PlayerActionDto
}

type SendPlayerIntoPortalCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	UnitId   uuid.UUID
}

type LeaveWorldCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
