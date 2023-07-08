package playerappsrv

import (
	"github.com/google/uuid"
)

type GetPlayerQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}

type GetPlayersQuery struct {
	WorldId  uuid.UUID
	PlayerId uuid.UUID
}
