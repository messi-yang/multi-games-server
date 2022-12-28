package livegamemodel

import "github.com/google/uuid"

type LiveGameId struct {
	id uuid.UUID
}

func NewLiveGameId(rawId string) (LiveGameId, error) {
	id, err := uuid.Parse(rawId)
	if err != nil {
		return LiveGameId{}, err
	}

	return LiveGameId{
		id: id,
	}, nil
}

func (liveGameId LiveGameId) ToString() string {
	return liveGameId.id.String()
}
