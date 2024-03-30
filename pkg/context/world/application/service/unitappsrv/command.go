package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type CreateStaticUnitCommand struct {
	Id        uuid.UUID
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

type RemoveStaticUnitCommand struct {
	Id uuid.UUID
}

type CreateFenceUnitCommand struct {
	Id        uuid.UUID
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

type RemoveFenceUnitCommand struct {
	Id uuid.UUID
}

type RotateUnitCommand struct {
	Id uuid.UUID
}
