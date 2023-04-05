package jsondto

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/google/uuid"
)

type UnitAggDto struct {
	ItemId    uuid.UUID     `json:"itemId"`
	Position  PositionVoDto `json:"position"`
	Direction int8          `json:"direction"`
}

func NewUnitAggDto(unit unitmodel.UnitAgg) UnitAggDto {
	return UnitAggDto{
		ItemId:    unit.GetItemId().Uuid(),
		Position:  NewPositionVoDto(unit.GetPosition()),
		Direction: unit.GetDirection().Int8(),
	}
}
