package worldaccessappsrv

import (
	"github.com/google/uuid"
)

type AddWorldMemberCommand struct {
	UserId  uuid.UUID
	WorldId uuid.UUID
	Role    string
}
