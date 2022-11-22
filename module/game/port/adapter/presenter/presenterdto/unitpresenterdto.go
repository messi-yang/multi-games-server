package presenterdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/module/game/domain/model/gamecommonmodel"
)

type UnitPresenterDto struct {
	Alive    bool                     `json:"alive"`
	ItemType gamecommonmodel.ItemType `json:"itemType"`
}

func NewUnitPresenterDto(unit gamecommonmodel.Unit) UnitPresenterDto {
	return UnitPresenterDto{
		Alive:    unit.GetAlive(),
		ItemType: unit.GetItemType(),
	}
}

func (dto UnitPresenterDto) ToValueObject() gamecommonmodel.Unit {
	return gamecommonmodel.NewUnit(dto.Alive, dto.ItemType)
}
