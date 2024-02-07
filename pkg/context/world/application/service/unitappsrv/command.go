package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type CreateStaticUnitCommand struct {
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

type RemoveStaticUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type CreateFenceUnitCommand struct {
	WorldId   uuid.UUID
	ItemId    uuid.UUID
	Position  dto.PositionDto
	Direction int8
}

type RemoveFenceUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type CreatePortalUnitCommand struct {
	WorldId        uuid.UUID
	ItemId         uuid.UUID
	Position       dto.PositionDto
	Direction      int8
	TargetPosition *dto.PositionDto
}

type RemovePortalUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type RotateUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}
