package userappsrv

import "github.com/google/uuid"

type GetUserByEmailAddressQuery struct {
	EmailAddress string
}

type GetUserQuery struct {
	UserId uuid.UUID
}

type GetUsersOfIdsQuery struct {
	UserIds []uuid.UUID
}
