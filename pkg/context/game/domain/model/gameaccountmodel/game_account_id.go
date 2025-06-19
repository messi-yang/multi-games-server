package gameaccountmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type GameAccountId struct {
	id uuid.UUID
}

// Interface Implementation Check
var _ domain.ValueObject[GameAccountId] = (*GameAccountId)(nil)

func NewGameAccountId(uuid uuid.UUID) GameAccountId {
	return GameAccountId{
		id: uuid,
	}
}

func (gameAccountId GameAccountId) IsEqual(otherGameAccountId GameAccountId) bool {
	return gameAccountId.id == otherGameAccountId.id
}

func (gameAccountId GameAccountId) Uuid() uuid.UUID {
	return gameAccountId.id
}
