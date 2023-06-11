package worldaccessappsrv

import "github.com/google/uuid"

type FindUserWorldRoleQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}

type GetUserWorldRolesQuery struct {
	WorldId uuid.UUID
}
