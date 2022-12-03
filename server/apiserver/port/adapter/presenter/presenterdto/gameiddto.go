package presenterdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/google/uuid"
)

type LiveGameIdPresenterDto []byte

func NewLiveGameIdPresenterDto(liveGameId livegamemodel.LiveGameId) LiveGameIdPresenterDto {
	return LiveGameIdPresenterDto(liveGameId.GetId().String())
}

func (dto LiveGameIdPresenterDto) ToValueObject() (livegamemodel.LiveGameId, error) {
	id, err := uuid.ParseBytes(dto)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return livegamemodel.NewLiveGameId(id), nil
}
