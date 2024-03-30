package portalunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type CreatePortalUnitCommand struct {
	Id             uuid.UUID
	WorldId        uuid.UUID
	ItemId         uuid.UUID
	Position       dto.PositionDto
	Direction      int8
	TargetPosition *dto.PositionDto
}

type RemovePortalUnitCommand struct {
	Id uuid.UUID
}
