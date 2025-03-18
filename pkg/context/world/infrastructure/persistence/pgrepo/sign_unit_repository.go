package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/signunitmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type signUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

// Interface Implementation Check
var _ signunitmodel.SignUnitRepo = (*signUnitRepo)(nil)

func NewSignUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository signunitmodel.SignUnitRepo) {
	return &signUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *signUnitRepo) Add(signUnit signunitmodel.SignUnit) error {
	unitModel := pgmodel.NewSignUnitModel(signUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(signUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *signUnitRepo) Get(id signunitmodel.SignUnitId) (unit signunitmodel.SignUnit, err error) {
	unitModel := pgmodel.UnitModel{}

	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumSign,
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}

	return pgmodel.ParseSignUnitModels(unitModel)
}

func (repo *signUnitRepo) Update(signUnit signunitmodel.SignUnit) error {
	unitModel := pgmodel.NewSignUnitModel(signUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(signUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			signUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}
		return transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			signUnit.GetWorldId().Uuid(),
			signUnit.GetPosition().GetX(),
			signUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumSign,
		).Select("*").Updates(unitModel).Error
	})
}

func (repo *signUnitRepo) Delete(signUnit signunitmodel.SignUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			signUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			signUnit.GetWorldId().Uuid(),
			signUnit.GetPosition().GetX(),
			signUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumSign,
		).Delete(&pgmodel.UnitModel{}).Error
	})
}
