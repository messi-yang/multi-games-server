package presenterdto

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/google/uuid"
)

type PlayerIdPresenterDto []byte

func NewPlayerIdPresenterDto(playerId gamecommonmodel.PlayerId) PlayerIdPresenterDto {
	return PlayerIdPresenterDto(playerId.GetId().String())
}

func (dto PlayerIdPresenterDto) ToValueObject() (gamecommonmodel.PlayerId, error) {
	id, err := uuid.ParseBytes(dto)
	if err != nil {
		return gamecommonmodel.PlayerId{}, err
	}
	return gamecommonmodel.NewPlayerId(id), nil
}
