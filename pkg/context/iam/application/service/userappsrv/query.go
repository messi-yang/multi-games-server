package userappsrv

import "github.com/google/uuid"

type FindUserByEmailAddressQuery struct {
	EmailAddress string
}

type GetUserQuery struct {
	UserId uuid.UUID
}
