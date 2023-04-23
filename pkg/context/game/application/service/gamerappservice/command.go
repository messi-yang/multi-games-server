package gamerappservice

import "github.com/google/uuid"

type CreateGamerCommand struct {
	UserId uuid.UUID
}
