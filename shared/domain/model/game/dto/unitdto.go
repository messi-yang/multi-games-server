package dto

import (
	commonValueobject "github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/common/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
)

type UnitDto struct {
	Alive    bool                       `json:"alive"`
	ItemType commonValueobject.ItemType `json:"itemType"`
}

func NewUnitDto(unit valueobject.Unit) UnitDto {
	return UnitDto{
		Alive:    unit.GetAlive(),
		ItemType: unit.GetItemType(),
	}
}

func (dto UnitDto) ToValueObject() valueobject.Unit {
	return valueobject.NewUnit(dto.Alive, dto.ItemType)
}
