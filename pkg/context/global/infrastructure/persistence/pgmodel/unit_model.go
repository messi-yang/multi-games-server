package pgmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
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
	Id           uuid.UUID    `gorm:"not null"`
	WorldId      uuid.UUID    `gorm:"not null"`
	PosX         int          `gorm:"not null"`
	PosZ         int          `gorm:"not null"`
	ItemId       uuid.UUID    `gorm:"not null"`
	Direction    int8         `gorm:"not null"`
	Label        *string      `gorm:""`
	Type         UnitTypeEnum `gorm:"not null"`
	InfoSnapshot pgtype.JSONB `gorm:"type:jsonb;not null"`
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
	return unitmodel.LoadUnit(
		unitmodel.NewUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		unitModel.Label,
		unitType,
		unitModel.InfoSnapshot,
	), nil
}

func NewEmbedUnitModel(embedUnit embedunitmodel.EmbedUnit) UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return UnitModel{
		WorldId:      embedUnit.GetWorldId().Uuid(),
		PosX:         embedUnit.GetPosition().GetX(),
		PosZ:         embedUnit.GetPosition().GetZ(),
		ItemId:       embedUnit.GetItemId().Uuid(),
		Direction:    embedUnit.GetDirection().Int8(),
		Label:        embedUnit.GetLabel(),
		Type:         UnitTypeEnumEmbed,
		Id:           embedUnit.GetId().Uuid(),
		InfoSnapshot: unitInfoSnapshotJsonb,
	}
}

func ParseEmbedUnitModels(unitModel UnitModel, embedUnitInfoModel EmbedUnitInfoModel) (unit embedunitmodel.EmbedUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(embedUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	embedCode, err := worldcommonmodel.NewEmbedCode(embedUnitInfoModel.EmbedCode)
	if err != nil {
		return unit, err
	}

	return embedunitmodel.LoadEmbedUnit(
		embedunitmodel.NewEmbedUnitId(embedUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		unitModel.Label,
		embedCode,
	), nil
}

func NewFenceUnitModel(embedUnit fenceunitmodel.FenceUnit) UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return UnitModel{
		Id:           embedUnit.GetId().Uuid(),
		WorldId:      embedUnit.GetWorldId().Uuid(),
		PosX:         embedUnit.GetPosition().GetX(),
		PosZ:         embedUnit.GetPosition().GetZ(),
		ItemId:       embedUnit.GetItemId().Uuid(),
		Direction:    embedUnit.GetDirection().Int8(),
		Type:         UnitTypeEnumFence,
		InfoSnapshot: unitInfoSnapshotJsonb,
	}
}

func ParseFenceUnitModels(unitModel UnitModel) (fenceunitmodel.FenceUnit, error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)

	return fenceunitmodel.LoadFenceUnit(
		fenceunitmodel.NewFenceUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
	), nil
}

func NewLinkUnitModel(linkUnit linkunitmodel.LinkUnit) UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return UnitModel{
		WorldId:      linkUnit.GetWorldId().Uuid(),
		PosX:         linkUnit.GetPosition().GetX(),
		PosZ:         linkUnit.GetPosition().GetZ(),
		ItemId:       linkUnit.GetItemId().Uuid(),
		Direction:    linkUnit.GetDirection().Int8(),
		Label:        linkUnit.GetLabel(),
		Type:         UnitTypeEnumLink,
		Id:           linkUnit.GetId().Uuid(),
		InfoSnapshot: unitInfoSnapshotJsonb,
	}
}

func ParseLinkUnitModels(unitModel UnitModel, linkUnitInfoModel LinkUnitInfoModel) (unit linkunitmodel.LinkUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(linkUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	url, err := globalcommonmodel.NewUrl(linkUnitInfoModel.Url)
	if err != nil {
		return unit, err
	}

	return linkunitmodel.LoadLinkUnit(
		linkunitmodel.NewLinkUnitId(linkUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		unitModel.Label,
		url,
	), nil
}

func NewPortalUnitModel(portalUnit portalunitmodel.PortalUnit) UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set(portalUnit.GetInfoSnapshot())

	return UnitModel{
		Id:           portalUnit.GetId().Uuid(),
		WorldId:      portalUnit.GetWorldId().Uuid(),
		PosX:         portalUnit.GetPosition().GetX(),
		PosZ:         portalUnit.GetPosition().GetZ(),
		ItemId:       portalUnit.GetItemId().Uuid(),
		Direction:    portalUnit.GetDirection().Int8(),
		Type:         UnitTypeEnumPortal,
		InfoSnapshot: unitInfoSnapshotJsonb,
	}
}

func ParsePortalUnitModels(unitModel UnitModel, portalUnitInfoModel PortalUnitInfoModel) (unit portalunitmodel.PortalUnit, err error) {
	worldId := globalcommonmodel.NewWorldId(portalUnitInfoModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	targetPosition := lo.TernaryF(
		portalUnitInfoModel.TargetPosX == nil,
		func() *worldcommonmodel.Position {
			return nil
		},
		func() *worldcommonmodel.Position {
			return commonutil.ToPointer(worldcommonmodel.NewPosition(*portalUnitInfoModel.TargetPosX, *portalUnitInfoModel.TargetPosZ))
		},
	)

	return portalunitmodel.LoadPortalUnit(
		portalunitmodel.NewPortalUnitId(portalUnitInfoModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
		targetPosition,
	), nil
}

func NewStaticUnitModel(staticUnit staticunitmodel.StaticUnit) UnitModel {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return UnitModel{
		Id:           staticUnit.GetId().Uuid(),
		WorldId:      staticUnit.GetWorldId().Uuid(),
		PosX:         staticUnit.GetPosition().GetX(),
		PosZ:         staticUnit.GetPosition().GetZ(),
		ItemId:       staticUnit.GetItemId().Uuid(),
		Direction:    staticUnit.GetDirection().Int8(),
		Type:         UnitTypeEnumStatic,
		InfoSnapshot: unitInfoSnapshotJsonb,
	}
}

func ParseStaticUnitModels(unitModel UnitModel) (staticunitmodel.StaticUnit, error) {
	worldId := globalcommonmodel.NewWorldId(unitModel.WorldId)
	pos := worldcommonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)

	return staticunitmodel.LoadStaticUnit(
		staticunitmodel.NewStaticUnitId(unitModel.Id),
		worldId,
		pos,
		worldcommonmodel.NewItemId(unitModel.ItemId),
		worldcommonmodel.NewDirection(unitModel.Direction),
	), nil
}
