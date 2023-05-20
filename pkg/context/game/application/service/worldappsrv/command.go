package worldappsrv

import (
	"github.com/google/uuid"
)

type CreateWorldCommand struct {
	UserId uuid.UUID
	Name   string
}

type UpdateWorldCommand struct {
	UserId  uuid.UUID
	WorldId uuid.UUID
	Name    string
}
