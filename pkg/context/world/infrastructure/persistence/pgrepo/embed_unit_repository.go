package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/embedunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelsFromEmbedUnit(embedUnit embedunitmodel.EmbedUnit) (pgmodel.EmbedUnitInfoModel, pgmodel.UnitModel, []pgmodel.OccupiedPositionModel) {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")

	return pgmodel.EmbedUnitInfoModel{
			Id:        embedUnit.GetId().Uuid(),
			WorldId:   embedUnit.GetWorldId().Uuid(),
			EmbedCode: embedUnit.GetEmbedCode().String(),
		},
		pgmodel.UnitModel{
			WorldId:      embedUnit.GetWorldId().Uuid(),
			PosX:         embedUnit.GetPosition().GetX(),
			PosZ:         embedUnit.GetPosition().GetZ(),
			ItemId:       embedUnit.GetItemId().Uuid(),
			Direction:    embedUnit.GetDirection().Int8(),
			Label:        embedUnit.GetLabel(),
			Type:         pgmodel.UnitTypeEnumEmbed,
			Id:           embedUnit.GetId().Uuid(),
			InfoSnapshot: unitInfoSnapshotJsonb,
		},
		pgmodel.NewOccupiedPositionsFromUnit(embedUnit.UnitEntity)
}

func parseModelsToEmbedUnit(unitModel pgmodel.UnitModel, embedUnitInfoModel pgmodel.EmbedUnitInfoModel) (unit embedunitmodel.EmbedUnit, err error) {
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

type embedUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

// Interface Implementation Check
var _ embedunitmodel.EmbedUnitRepo = (*embedUnitRepo)(nil)

func NewEmbedUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository embedunitmodel.EmbedUnitRepo) {
	return &embedUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *embedUnitRepo) Add(embedUnit embedunitmodel.EmbedUnit) error {
	embedUnitInfoModel, unitModel, occupiedPositionModels := newModelsFromEmbedUnit(embedUnit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&embedUnitInfoModel).Error; err != nil {
			return err
		}
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *embedUnitRepo) Get(id embedunitmodel.EmbedUnitId) (unit embedunitmodel.EmbedUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	embedUnitInfoModel := pgmodel.EmbedUnitInfoModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumEmbed,
		).First(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&embedUnitInfoModel).Error
	}); err != nil {
		return unit, err
	}

	return parseModelsToEmbedUnit(unitModel, embedUnitInfoModel)
}

func (repo *embedUnitRepo) Update(embedUnit embedunitmodel.EmbedUnit) error {
	embedUnitInfoModel, _, unitModel := newModelsFromEmbedUnit(embedUnit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			embedUnit.GetWorldId().Uuid(),
			embedUnit.GetPosition().GetX(),
			embedUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumEmbed,
		).Select("*").Updates(unitModel).Error; err != nil {
			return err
		}
		return transaction.Model(&pgmodel.EmbedUnitInfoModel{}).Where(
			"id = ?",
			embedUnit.GetId().Uuid(),
		).Select("*").Updates(embedUnitInfoModel).Error
	})
}

func (repo *embedUnitRepo) Delete(embedUnit embedunitmodel.EmbedUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			embedUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			embedUnit.GetWorldId().Uuid(),
			embedUnit.GetPosition().GetX(),
			embedUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumEmbed,
		).Delete(&pgmodel.UnitModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			embedUnit.GetId().Uuid(),
		).Delete(&pgmodel.EmbedUnitInfoModel{}).Error
	})
}
