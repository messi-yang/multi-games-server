package worldappsrv

import (
	"github.com/google/uuid"
)

type CreateWorldCommand struct {
	UserId uuid.UUID
	Name   string
}

type UpdateWorldCommand struct {
	WorldId uuid.UUID
	Name    string
}

type DeleteWorldCommand struct {
	WorldId uuid.UUID
}
