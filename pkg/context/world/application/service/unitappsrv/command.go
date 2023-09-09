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

type CreatePortalUnitCommand struct {
	WorldId        uuid.UUID
	ItemId         uuid.UUID
	Position       dto.PositionDto
	Direction      int8
	TargetPosition *dto.PositionDto
}

type RotateUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type RemoveUnitCommand struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}
