package playermodel

import "github.com/google/uuid"

type PlayerIdVo struct {
	id uuid.UUID
}

func ParsePlayerIdVo(uuidStr string) (PlayerIdVo, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return PlayerIdVo{}, err
	}

	return PlayerIdVo{
		id: id,
	}, nil
}

func (playerId PlayerIdVo) IsEqual(otherPlayerId PlayerIdVo) bool {
	return playerId.String() == otherPlayerId.String()
}

func (playerId PlayerIdVo) String() string {
	return playerId.id.String()
}
