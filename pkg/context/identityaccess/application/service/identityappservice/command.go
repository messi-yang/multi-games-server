package identityappservice

import "github.com/google/uuid"

type LoginOrRegisterCommand struct {
	EmailAddress string
}

type GenerateAccessTokenCommand struct {
	UserId uuid.UUID
}
