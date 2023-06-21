package gameappsrv

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/application/dto"
	"github.com/google/uuid"
)

type GetUnitQuery struct {
	WorldId  uuid.UUID
	Position dto.PositionDto
}

type GetPlayerQuery struct {
	PlayerId uuid.UUID
}

type GetNearbyPlayersQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type GetNearbyUnitsQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
