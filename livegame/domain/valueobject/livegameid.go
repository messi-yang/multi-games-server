package valueobject

import "github.com/google/uuid"

type LiveGameId struct {
	id uuid.UUID
}

func NewLiveGameId(id uuid.UUID) LiveGameId {
	return LiveGameId{
		id: id,
	}
}

func (liveGameId LiveGameId) GetId() uuid.UUID {
	return liveGameId.id
}
