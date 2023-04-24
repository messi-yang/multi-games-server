package gameappsrv

import (
	"github.com/google/uuid"
)

type ChangeHeldItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	ItemId   uuid.UUID
}

type PlaceItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type RemoveItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type EnterWorldCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
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
