package worldmemberappsrv

import "github.com/google/uuid"

type GetWorldMemberOfUserQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}
