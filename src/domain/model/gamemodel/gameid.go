package gamemodel

import "github.com/google/uuid"

type GameId struct {
	id uuid.UUID
}

func NewGameId(uuidStr string) (GameId, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return GameId{}, err
	}

	return GameId{
		id: id,
	}, nil
}

func (gameId GameId) ToString() string {
	return gameId.id.String()
}
