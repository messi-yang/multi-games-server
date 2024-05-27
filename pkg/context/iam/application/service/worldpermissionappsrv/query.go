package worldpermissionappsrv

import "github.com/google/uuid"

type CanUpdateWorldQuery struct {
	WorldId uuid.UUID
	UserId  uuid.UUID
}
