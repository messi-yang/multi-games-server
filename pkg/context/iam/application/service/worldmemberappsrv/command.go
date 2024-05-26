package worldmemberappsrv

import (
	"github.com/google/uuid"
)

type DeleteAllWorldMembersInWorldCommand struct {
	WorldId uuid.UUID
}
