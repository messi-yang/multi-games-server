package worldmemberappsrv

import (
	"github.com/google/uuid"
)

type AddWorldMemberCommand struct {
	UserId  uuid.UUID
	WorldId uuid.UUID
	Role    string
}

type DeleteAllWorldMembersInWorldCommand struct {
	WorldId uuid.UUID
}
