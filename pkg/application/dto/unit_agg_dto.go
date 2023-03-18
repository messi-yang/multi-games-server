package dto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
)

type UnitVoDto struct {
	ItemId    string        `json:"itemId"`
	Position  PositionVoDto `json:"position"`
	Direction int8          `json:"direction"`
}

func NewUnitVoDto(unit unitmodel.UnitAgg) UnitVoDto {
	return UnitVoDto{
		ItemId:    unit.GetItemId().String(),
		Position:  NewPositionVoDto(unit.GetPosition()),
		Direction: unit.GetDirection().Int8(),
	}
}
