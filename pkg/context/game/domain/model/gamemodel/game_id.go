package gamemodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type GameId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[GameId] = (*GameId)(nil)

func NewGameId(id uuid.UUID) GameId {
	return GameId{
		id: id,
	}
}

func (gameId GameId) IsEqual(otherGameId GameId) bool {
	return gameId.id == otherGameId.id
}

func (gameId GameId) Uuid() uuid.UUID {
	return gameId.id
}
