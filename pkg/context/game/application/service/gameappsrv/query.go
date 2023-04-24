package gameappsrv

import (
	"github.com/google/uuid"
)

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
