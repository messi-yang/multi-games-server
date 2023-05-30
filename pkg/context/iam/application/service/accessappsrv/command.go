package accessappsrv

import (
	"github.com/google/uuid"
)

type AssignUserToWorldRoleCommand struct {
	UserId        uuid.UUID
	WorldId       uuid.UUID
	WorldRoleName string
}
