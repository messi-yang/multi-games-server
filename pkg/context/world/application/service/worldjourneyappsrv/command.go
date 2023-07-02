package worldjourneyappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type ChangeHeldItemCommand struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
	ItemId   uuid.UUID
}

type PlaceUnitCommand struct {
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

type RemoveUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type EnterWorldCommand struct {
	WorldId uuid.UUID
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
