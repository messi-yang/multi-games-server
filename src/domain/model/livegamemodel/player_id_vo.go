package livegamemodel

import "github.com/google/uuid"

type PlayerIdVo struct {
	id uuid.UUID
}

func NewPlayerIdVo(uuidStr string) (PlayerIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return PlayerIdVo{}, err
	}

	return PlayerIdVo{
		id: id,
	}, nil
}

func (playerId PlayerIdVo) isEqual(otherPlayerId PlayerIdVo) bool {
	return playerId.ToString() == otherPlayerId.ToString()
}

func (playerId PlayerIdVo) ToString() string {
	return playerId.id.String()
}
