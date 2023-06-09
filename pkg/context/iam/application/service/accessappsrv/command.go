package accessappsrv

import (
	"github.com/google/uuid"
)

type AssignWorldRoleToUserCommand struct {
	UserId    uuid.UUID
	WorldId   uuid.UUID
	WorldRole string
}
