package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UnitDto struct {
	Id        uuid.UUID    `json:"id"`
	ItemId    uuid.UUID    `json:"itemId"`
	Position  PositionDto  `json:"position"`
	Direction int8         `json:"direction"`
	Dimension DimensionDto `json:"dimension"`
	Label     *string      `json:"label"`
	Color     *string      `json:"color"`
	Type      string       `json:"type"`
}

func NewUnitDto(unit unitmodel.Unit) UnitDto {
	color := lo.TernaryF(
		unit.GetColor() == nil,
		func() *string { return nil },
		func() *string { return commonutil.ToPointer(unit.GetColor().HexString()) },
	)
	return UnitDto{
		Id:        unit.GetId().Uuid(),
		ItemId:    unit.GetItemId().Uuid(),
		Position:  NewPositionDto(unit.GetPosition()),
		Direction: unit.GetDirection().Int8(),
		Dimension: NewDimensionDto(unit.GetDimension()),
		Label:     unit.GetLabel(),
		Color:     color,
		Type:      unit.GetType().String(),
	}
}
