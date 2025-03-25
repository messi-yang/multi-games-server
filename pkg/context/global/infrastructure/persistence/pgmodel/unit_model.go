package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/colorunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/signunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UnitModel struct {
	Id             uuid.UUID    `gorm:"not null"`
	WorldId        uuid.UUID    `gorm:"not null"`
	PosX           int          `gorm:"not null"`
	PosZ           int          `gorm:"not null"`
	ItemId         uuid.UUID    `gorm:"not null"`
	Direction      int8         `gorm:"not null"`
	DimensionWidth int          `gorm:"not null"`
	DimensionDepth int          `gorm:"not null"`
	Label          *string      `gorm:""`
	Color          *string      `gorm:""`
	Type           UnitTypeEnum `gorm:"not null"`
}

func (UnitModel) TableName() string {
	return "units"
}

func NewUnitModel(unit unitmodel.Unit) UnitModel {
	color := unit.GetColor()
	colorString := lo.TernaryF(
		color == nil,
		func() *string { return nil },
		func() *string {
			return commonutil.ToPointer(color.HexString())
		},
	)

	return UnitModel{
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Label:          unit.GetLabel(),
		Color:          colorString,
		Type:           UnitTypeEnum(unit.GetType().String()),
		Id:             unit.GetId().Uuid(),
	}
}

func ParseUnitModel(unitModel UnitModel) (unit unitmodel.Unit, err error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	unitType, err := worldcommonmodel.NewUnitType(string(unitModel.Type))
	if err != nil {
		return unit, err
	}
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}
	color := lo.TernaryF(
		unitModel.Color == nil,
		func() *globalcommonmodel.Color { return nil },
		func() *globalcommonmodel.Color {
			color, err := globalcommonmodel.NewColorFromHexString(*unitModel.Color)
			if err != nil {
				return nil
			}
			return commonutil.ToPointer(color)
		},
	)

	return unitmodel.LoadUnit(
		unitmodel.NewUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
		unitModel.Label,
		color,
		unitType,
	), nil
}

func NewEmbedUnitModel(unit embedunitmodel.EmbedUnit) UnitModel {
	return UnitModel{
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Label:          unit.GetLabel(),
		Type:           UnitTypeEnumEmbed,
		Id:             unit.GetId().Uuid(),
	}
}

func ParseEmbedUnitModels(unitModel UnitModel, embedUnitInfoModel EmbedUnitInfoModel) (unit embedunitmodel.EmbedUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(embedUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	embedCode, err := worldcommonmodel.NewEmbedCode(embedUnitInfoModel.EmbedCode)
	if err != nil {
		return unit, err
	}
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return embedunitmodel.LoadEmbedUnit(
		embedunitmodel.NewEmbedUnitId(embedUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
		unitModel.Label,
		embedCode,
	), nil
}

func NewFenceUnitModel(unit fenceunitmodel.FenceUnit) UnitModel {
	return UnitModel{
		Id:             unit.GetId().Uuid(),
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Type:           UnitTypeEnumFence,
	}
}

func ParseFenceUnitModels(unitModel UnitModel) (unit fenceunitmodel.FenceUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return fenceunitmodel.LoadFenceUnit(
		fenceunitmodel.NewFenceUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
	), nil
}

func NewLinkUnitModel(unit linkunitmodel.LinkUnit) UnitModel {
	return UnitModel{
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Label:          unit.GetLabel(),
		Type:           UnitTypeEnumLink,
		Id:             unit.GetId().Uuid(),
	}
}

func ParseLinkUnitModels(unitModel UnitModel, linkUnitInfoModel LinkUnitInfoModel) (unit linkunitmodel.LinkUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(linkUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	url, err := globalcommonmodel.NewUrl(linkUnitInfoModel.Url)
	if err != nil {
		return unit, err
	}
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return linkunitmodel.LoadLinkUnit(
		linkunitmodel.NewLinkUnitId(linkUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
		unitModel.Label,
		url,
	), nil
}

func NewPortalUnitModel(unit portalunitmodel.PortalUnit) UnitModel {
	return UnitModel{
		Id:             unit.GetId().Uuid(),
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Type:           UnitTypeEnumPortal,
	}
}

func ParsePortalUnitModels(unitModel UnitModel, portalUnitInfoModel PortalUnitInfoModel) (unit portalunitmodel.PortalUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(portalUnitInfoModel.WorldId)
	targetUnitId := lo.TernaryF(
		portalUnitInfoModel.TargetUnitId == nil,
		func() *portalunitmodel.PortalUnitId { return nil },
		func() *portalunitmodel.PortalUnitId {
			return commonutil.ToPointer(portalunitmodel.NewPortalUnitId(*portalUnitInfoModel.TargetUnitId))
		},
	)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return portalunitmodel.LoadPortalUnit(
		portalunitmodel.NewPortalUnitId(portalUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
		targetUnitId,
	), nil
}

func NewStaticUnitModel(unit staticunitmodel.StaticUnit) UnitModel {
	return UnitModel{
		Id:             unit.GetId().Uuid(),
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Type:           UnitTypeEnumStatic,
	}
}

func ParseStaticUnitModels(unitModel UnitModel) (unit staticunitmodel.StaticUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return staticunitmodel.LoadStaticUnit(
		staticunitmodel.NewStaticUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
	), nil
}

func NewColorUnitModel(unit colorunitmodel.ColorUnit) UnitModel {
	return UnitModel{
		Id:             unit.GetId().Uuid(),
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Label:          unit.GetLabel(),
		Color:          commonutil.ToPointer(unit.GetColor().HexString()),
		Type:           UnitTypeEnumColor,
	}
}

func ParseColorUnitModels(unitModel UnitModel) (unit colorunitmodel.ColorUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	color, err := globalcommonmodel.NewColorFromHexString(*unitModel.Color)
	if err != nil {
		return unit, err
	}
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return colorunitmodel.LoadColorUnit(
		colorunitmodel.NewColorUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
		unitModel.Label,
		commonutil.ToPointer(color),
	), nil
}

func NewSignUnitModel(unit signunitmodel.SignUnit) UnitModel {
	return UnitModel{
		Id:             unit.GetId().Uuid(),
		WorldId:        unit.GetWorldId().Uuid(),
		PosX:           unit.GetPosition().GetX(),
		PosZ:           unit.GetPosition().GetZ(),
		ItemId:         unit.GetItemId().Uuid(),
		Direction:      unit.GetDirection().Int8(),
		DimensionWidth: unit.GetDimension().GetWidth(),
		DimensionDepth: unit.GetDimension().GetDepth(),
		Label:          commonutil.ToPointer(unit.GetLabel()),
		Type:           UnitTypeEnumSign,
	}
}

func ParseSignUnitModels(unitModel UnitModel) (unit signunitmodel.SignUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	dimension, err := worldcommonmodel.NewDimension(unitModel.DimensionWidth, unitModel.DimensionDepth)
	if err != nil {
		return unit, err
	}

	return signunitmodel.LoadSignUnit(
		signunitmodel.NewSignUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		dimension,
		*unitModel.Label,
	), nil
}
