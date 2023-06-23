package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel/unitmodel"
	"github.com/google/uuid"
)

type UnitDto struct {
	ItemId    uuid.UUID   `json:"itemId"`
	Position  PositionDto `json:"position"`
	Direction int8        `json:"direction"`
}

func NewUnitDto(unit unitmodel.Unit) UnitDto {
	return UnitDto{
		ItemId:    unit.GetItemId().Uuid(),
		Position:  NewPositionDto(unit.GetPosition()),
		Direction: unit.GetDirection().Int8(),
	}
}
