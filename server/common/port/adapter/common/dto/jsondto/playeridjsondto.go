package jsondto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/google/uuid"
)

type PlayerIdJsonDto []byte

func NewPlayerIdJsonDto(playerId gamecommonmodel.PlayerId) PlayerIdJsonDto {
	return PlayerIdJsonDto(playerId.GetId().String())
}

func (dto PlayerIdJsonDto) ToValueObject() (gamecommonmodel.PlayerId, error) {
	id, err := uuid.ParseBytes(dto)
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return gamecommonmodel.NewPlayerId(id), nil
}
