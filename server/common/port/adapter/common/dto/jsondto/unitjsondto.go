package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/domainmodel/itemmodel"
)

type UnitJsonDto struct {
	Alive  bool   `json:"alive"`
	ItemId string `json:"itemId"`
}

func NewUnitJsonDto(unit commonmodel.Unit) UnitJsonDto {
	return UnitJsonDto{
		Alive:  unit.GetAlive(),
		ItemId: unit.GetItemId().ToString(),
	}
}

func (dto UnitJsonDto) ToValueObject() (commonmodel.Unit, error) {
	itemId, err := itemmodel.NewItemId(dto.ItemId)
	if err != nil {
		return commonmodel.Unit{}, err
	}
	return commonmodel.NewUnit(itemId), nil
}
