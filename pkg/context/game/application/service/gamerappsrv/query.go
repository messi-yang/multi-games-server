package gamerappsrv

import "github.com/google/uuid"

type GetGamerByUserIdQuery struct {
	UserId uuid.UUID
}

type QueryGamersQuery struct {
}
