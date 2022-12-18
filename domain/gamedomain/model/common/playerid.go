package common

import "github.com/google/uuid"

type PlayerId struct {
	id uuid.UUID
}

func NewPlayerId(rawId string) (PlayerId, error) {
	id, err := uuid.Parse(rawId)
	if err != nil {
		return PlayerId{}, err
	}

	return PlayerId{
		id: id,
	}, nil
}

func (playerId PlayerId) ToString() string {
	return playerId.id.String()
}
