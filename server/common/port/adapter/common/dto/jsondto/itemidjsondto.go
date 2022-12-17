package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/itemmodel"
	"github.com/google/uuid"
)

type ItemIdJsonDto []byte

func NewItemIdJsonDto(itemId itemmodel.ItemId) ItemIdJsonDto {
	return ItemIdJsonDto(itemId.GetId().String())
}

func (dto ItemIdJsonDto) ToValueObject() (itemmodel.ItemId, error) {
	id, err := uuid.ParseBytes(dto)
	if err != nil {
		return itemmodel.ItemId{}, err
	}
	return itemmodel.NewItemId(id), nil
}
