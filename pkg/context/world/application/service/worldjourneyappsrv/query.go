package worldjourneyappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/google/uuid"
)

type GetUnitQuery struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type GetPlayerQuery struct {
	PlayerId uuid.UUID
}

type GetPlayersQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type GetUnitsQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
