package valueobject

import "github.com/google/uuid"

type GameId struct {
	id uuid.UUID
}

func NewGameId(id uuid.UUID) GameId {
	return GameId{
		id: id,
	}
}

func (gameId GameId) GetId() uuid.UUID {
	return gameId.id
}
