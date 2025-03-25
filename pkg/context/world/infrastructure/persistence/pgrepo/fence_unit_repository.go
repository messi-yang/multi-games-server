package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/fenceunitmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

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
	unitModel := pgmodel.NewFenceUnitModel(fenceUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(fenceUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *fenceUnitRepo) Update(fenceUnit fenceunitmodel.FenceUnit) error {
	unitModel := pgmodel.NewFenceUnitModel(fenceUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(fenceUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			fenceUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}
		return transaction.Save(&unitModel).Error
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

	return pgmodel.ParseFenceUnitModels(unitModel)
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
			"id = ?",
			fenceUnit.GetId().Uuid(),
		).Delete(&pgmodel.UnitModel{}).Error
	})
}
