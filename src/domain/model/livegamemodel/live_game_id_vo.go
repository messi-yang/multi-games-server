package livegamemodel

import "github.com/google/uuid"

type LiveGameIdVo struct {
	id uuid.UUID
}

func NewLiveGameIdVo(uuidStr string) (LiveGameIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return LiveGameIdVo{}, err
	}

	return LiveGameIdVo{
		id: id,
	}, nil
}

func (liveGameId LiveGameIdVo) ToString() string {
	return liveGameId.id.String()
}
