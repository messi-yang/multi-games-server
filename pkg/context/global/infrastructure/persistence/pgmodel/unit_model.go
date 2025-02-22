package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/colorunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
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
	InfoSnapshot   pgtype.JSONB `gorm:"type:jsonb;not null"`
}

func (UnitModel) TableName() string {
	return "units"
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
		unitModel.InfoSnapshot,
	), nil
}

func NewEmbedUnitModel(unit embedunitmodel.EmbedUnit) UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

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
		InfoSnapshot:   unitInfoSnapshotJsonb,
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
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

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
		InfoSnapshot:   unitInfoSnapshotJsonb,
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
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

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
		InfoSnapshot:   unitInfoSnapshotJsonb,
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
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set(unit.GetInfoSnapshot())

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
		InfoSnapshot:   unitInfoSnapshotJsonb,
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
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

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
		InfoSnapshot:   unitInfoSnapshotJsonb,
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
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return UnitModel{
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
		Id:             unit.GetId().Uuid(),
		InfoSnapshot:   unitInfoSnapshotJsonb,
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
