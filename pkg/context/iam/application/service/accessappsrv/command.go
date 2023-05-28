package accessappsrv

import (
	"github.com/google/uuid"
)

type AssignWorldRoleCommand struct {
	UserId        uuid.UUID
	WorldId       uuid.UUID
	WorldRoleName string
}
