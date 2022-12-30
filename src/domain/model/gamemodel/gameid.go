package gamemodel

import "github.com/google/uuid"

type GameId struct {
	id uuid.UUID
}

func NewGameId(rawId string) (GameId, error) {
	id, err := uuid.Parse(rawId)
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
