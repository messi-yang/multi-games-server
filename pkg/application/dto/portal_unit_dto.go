package dto

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type PortalUnitDto struct {
	Id           uuid.UUID    `json:"id"`
	ItemId       uuid.UUID    `json:"itemId"`
	Position     PositionDto  `json:"position"`
	Direction    int8         `json:"direction"`
	Dimension    DimensionDto `json:"dimension"`
	Label        *string      `json:"label"`
	Color        *string      `json:"color"`
	Type         string       `json:"type"`
	TargetUnitId *uuid.UUID   `json:"targetUnitId"`
}

func NewPortalUnitDto(portalUnit portalunitmodel.PortalUnit) PortalUnitDto {
	color := lo.TernaryF(
		portalUnit.GetColor() == nil,
		func() *string { return nil },
		func() *string { return commonutil.ToPointer(portalUnit.GetColor().HexString()) },
	)
	return PortalUnitDto{
		Id:           portalUnit.GetId().Uuid(),
		ItemId:       portalUnit.GetItemId().Uuid(),
		Position:     NewPositionDto(portalUnit.GetPosition()),
		Direction:    portalUnit.GetDirection().Int8(),
		Dimension:    NewDimensionDto(portalUnit.GetDimension()),
		Label:        portalUnit.GetLabel(),
		Color:        color,
		Type:         portalUnit.GetType().String(),
		TargetUnitId: commonutil.ToPointer(portalUnit.GetTargetUnitId().Uuid()),
	}
}
