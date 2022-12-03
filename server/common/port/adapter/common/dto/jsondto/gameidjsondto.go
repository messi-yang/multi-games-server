package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/google/uuid"
)

type LiveGameIdJsonDto []byte

func NewLiveGameIdJsonDto(liveGameId livegamemodel.LiveGameId) LiveGameIdJsonDto {
	return LiveGameIdJsonDto(liveGameId.GetId().String())
}

func (dto LiveGameIdJsonDto) ToValueObject() (livegamemodel.LiveGameId, error) {
	id, err := uuid.ParseBytes(dto)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	return livegamemodel.NewLiveGameId(id), nil
}
