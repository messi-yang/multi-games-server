package linkunitappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type GetLinkUnitUrlQuery struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}
