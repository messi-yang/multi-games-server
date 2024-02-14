package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/google/uuid"
)

type UnitDto struct {
	Id        uuid.UUID   `json:"id"`
	ItemId    uuid.UUID   `json:"itemId"`
	Position  PositionDto `json:"position"`
	Direction int8        `json:"direction"`
	Label     *string     `json:"label"`
	Type      string      `json:"type"`
	Info      any         `json:"info"`
}

func NewUnitDto(unit unitmodel.Unit) UnitDto {
	return UnitDto{
		Id:        unit.GetId().Uuid(),
		ItemId:    unit.GetItemId().Uuid(),
		Position:  NewPositionDto(unit.GetPosition()),
		Direction: unit.GetDirection().Int8(),
		Label:     unit.GetLabel(),
		Type:      unit.GetType().String(),
		Info:      unit.GetInfo(),
	}
}
