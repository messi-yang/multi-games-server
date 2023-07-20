package worldpermissionappsrv

import "github.com/google/uuid"

type CanGetWorldMembersQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}

type CanUpdateWorldQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}

type CanDeleteWorldQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}
