package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/jackc/pgtype"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelFromFenceUnit(fenceUnit fenceunitmodel.FenceUnit) (pgmodel.UnitModel, []pgmodel.OccupiedPositionModel) {
	unitInfoSnapshotJsonb := pgtype.JSONB{}
	unitInfoSnapshotJsonb.Set("null")
	return pgmodel.UnitModel{
		Id:           fenceUnit.GetId().Uuid(),
		WorldId:      fenceUnit.GetWorldId().Uuid(),
		PosX:         fenceUnit.GetPosition().GetX(),
		PosZ:         fenceUnit.GetPosition().GetZ(),
		ItemId:       fenceUnit.GetItemId().Uuid(),
		Direction:    fenceUnit.GetDirection().Int8(),
		Type:         pgmodel.UnitTypeEnumFence,
		InfoSnapshot: unitInfoSnapshotJsonb,
	}, pgmodel.NewOccupiedPositionsFromUnit(fenceUnit.UnitEntity)
}

func parseModelToFenceUnit(unitModel pgmodel.UnitModel) (fenceunitmodel.FenceUnit, error) {
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

type fenceUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewFenceUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository fenceunitmodel.FenceUnitRepo) {
	return &fenceUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *fenceUnitRepo) Add(fenceUnit fenceunitmodel.FenceUnit) error {
	unitModel, occupiedPositionModels := newModelFromFenceUnit(fenceUnit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *fenceUnitRepo) Update(fenceUnit fenceunitmodel.FenceUnit) error {
	unitModel, _ := newModelFromFenceUnit(fenceUnit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			unitModel.WorldId,
			unitModel.PosX,
			unitModel.PosZ,
			pgmodel.UnitTypeEnumFence,
		).Select("*").Updates(unitModel).Error
	})
}

func (repo *fenceUnitRepo) Get(id fenceunitmodel.FenceUnitId) (unit fenceunitmodel.FenceUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumFence,
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}

	return parseModelToFenceUnit(unitModel)
}

func (repo *fenceUnitRepo) Delete(fenceUnit fenceunitmodel.FenceUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			fenceUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			fenceUnit.GetWorldId().Uuid(),
			fenceUnit.GetPosition().GetX(),
			fenceUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumFence,
		).Delete(&pgmodel.UnitModel{}).Error
	})
}
