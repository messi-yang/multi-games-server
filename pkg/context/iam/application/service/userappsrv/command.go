package userappsrv

import "github.com/google/uuid"

type UpdateUserCommand struct {
	UserId       uuid.UUID
	Username     string
	FriendlyName string
}
