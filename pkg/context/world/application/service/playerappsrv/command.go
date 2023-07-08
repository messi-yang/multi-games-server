package playerappsrv

import (
	"github.com/google/uuid"
)

type ChangeHeldItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	ItemId   uuid.UUID
}

type EnterWorldCommand struct {
	WorldId          uuid.UUID
	PlayerHeldItemId uuid.UUID
}

type MoveCommand struct {
	WorldId   uuid.UUID
	PlayerId  uuid.UUID
	Direction int8
}

type LeaveWorldCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
