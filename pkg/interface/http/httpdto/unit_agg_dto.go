package httpdto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/google/uuid"
)

type UnitVoDto struct {
	ItemId    uuid.UUID     `json:"itemId"`
	Position  PositionVoDto `json:"position"`
	Direction int8          `json:"direction"`
}

func NewUnitVoDto(unit unitmodel.UnitAgg) UnitVoDto {
	return UnitVoDto{
		ItemId:    unit.GetItemId().Uuid(),
		Position:  NewPositionVoDto(unit.GetPosition()),
		Direction: unit.GetDirection().Int8(),
	}
}
