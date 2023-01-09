package livegamemodel

import "github.com/google/uuid"

type LiveGameId struct {
	id uuid.UUID
}

func NewLiveGameId(uuidStr string) (LiveGameId, error) {
	id, err := uuid.Parse(uuidStr)
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
