package playermodel

import "github.com/google/uuid"

type PlayerId struct {
	id uuid.UUID
}

func NewPlayerId(uuidStr string) (PlayerId, error) {
	id, err := uuid.Parse(uuidStr)
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
