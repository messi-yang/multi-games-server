package worldappsrv

import (
	"github.com/google/uuid"
)

type CreateWorldCommand struct {
	GamerId uuid.UUID
}
