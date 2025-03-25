package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/colorunitmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type colorUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

// Interface Implementation Check
var _ colorunitmodel.ColorUnitRepo = (*colorUnitRepo)(nil)

func NewColorUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository colorunitmodel.ColorUnitRepo) {
	return &colorUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *colorUnitRepo) Add(colorUnit colorunitmodel.ColorUnit) error {
	unitModel := pgmodel.NewColorUnitModel(colorUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(colorUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *colorUnitRepo) Get(id colorunitmodel.ColorUnitId) (unit colorunitmodel.ColorUnit, err error) {
	unitModel := pgmodel.UnitModel{}

	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumColor,
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}

	return pgmodel.ParseColorUnitModels(unitModel)
}

func (repo *colorUnitRepo) Update(colorUnit colorunitmodel.ColorUnit) error {
	unitModel := pgmodel.NewColorUnitModel(colorUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(colorUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			colorUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}
		return transaction.Save(&unitModel).Error
	})
}

func (repo *colorUnitRepo) Delete(colorUnit colorunitmodel.ColorUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			colorUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			colorUnit.GetId().Uuid(),
		).Delete(&pgmodel.UnitModel{}).Error
	})
}
