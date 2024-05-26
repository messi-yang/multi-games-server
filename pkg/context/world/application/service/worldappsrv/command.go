package worldappsrv

import (
	"github.com/google/uuid"
)

type UpdateWorldCommand struct {
	WorldId uuid.UUID
	Name    string
}

type DeleteWorldCommand struct {
	WorldId uuid.UUID
}
