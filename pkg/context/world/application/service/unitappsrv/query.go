package unitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type GetUnitQuery struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type GetUnitsQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
