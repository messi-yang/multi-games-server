package identityappsrv

import "github.com/google/uuid"

type LoginOrRegisterCommand struct {
	EmailAddress string
}

type LoginCommand struct {
	UserId uuid.UUID
}

type RegisterCommand struct {
	EmailAddress string
}

type GenerateAccessTokenCommand struct {
	UserId uuid.UUID
}
