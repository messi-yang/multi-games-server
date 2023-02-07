package gamemodel

import "github.com/google/uuid"

type GameIdVo struct {
	id uuid.UUID
}

func NewGameIdVo(uuidStr string) (GameIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return GameIdVo{}, err
	}

	return GameIdVo{
		id: id,
	}, nil
}

func (gameId GameIdVo) ToString() string {
	return gameId.id.String()
}
