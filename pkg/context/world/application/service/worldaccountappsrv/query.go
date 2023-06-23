package worldaccountappsrv

import "github.com/google/uuid"

type GetWorldAccountOfUserQuery struct {
	UserId uuid.UUID
}

type QueryWorldAccountsQuery struct {
}
