package worldaccountappsrv

import "github.com/google/uuid"

type CreateWorldAccountCommand struct {
	UserId uuid.UUID
}
