package worldappsrv

import (
	"github.com/google/uuid"
)

type CreateWorldCommand struct {
	GamerId uuid.UUID
}

type UpdateWorldCommand struct {
	GamerId uuid.UUID
	WorldId uuid.UUID
	Name    string
}
