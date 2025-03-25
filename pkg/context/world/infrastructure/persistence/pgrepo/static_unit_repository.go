package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/staticunitmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type staticUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewStaticUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository staticunitmodel.StaticUnitRepo) {
	return &staticUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *staticUnitRepo) Add(staticUnit staticunitmodel.StaticUnit) error {
	unitModel := pgmodel.NewStaticUnitModel(staticUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(staticUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *staticUnitRepo) Update(staticUnit staticunitmodel.StaticUnit) error {
	unitModel := pgmodel.NewStaticUnitModel(staticUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(staticUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			staticUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}
		return transaction.Save(&unitModel).Error
	})
}

func (repo *staticUnitRepo) Get(id staticunitmodel.StaticUnitId) (unit staticunitmodel.StaticUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumStatic,
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}

	return pgmodel.ParseStaticUnitModels(unitModel)
}

func (repo *staticUnitRepo) Delete(staticUnit staticunitmodel.StaticUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			staticUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			staticUnit.GetId().Uuid(),
		).Delete(&pgmodel.UnitModel{}).Error
	})
}
