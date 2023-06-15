package worldaccessappsrv

import "github.com/google/uuid"

type FindWorldMemberQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}

type GetWorldMembersQuery struct {
	WorldId uuid.UUID
}
