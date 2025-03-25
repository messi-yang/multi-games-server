package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/linkunitmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type linkUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

// Interface Implementation Check
var _ linkunitmodel.LinkUnitRepo = (*linkUnitRepo)(nil)

func NewLinkUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository linkunitmodel.LinkUnitRepo) {
	return &linkUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *linkUnitRepo) Add(linkUnit linkunitmodel.LinkUnit) error {
	linkUnitInfoModel := pgmodel.NewLinkUnitInfoModel(linkUnit)
	unitModel := pgmodel.NewLinkUnitModel(linkUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(linkUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&linkUnitInfoModel).Error; err != nil {
			return err
		}
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *linkUnitRepo) Get(id linkunitmodel.LinkUnitId) (unit linkunitmodel.LinkUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	linkUnitInfoModel := pgmodel.LinkUnitInfoModel{}

	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumLink,
		).First(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&linkUnitInfoModel).Error
	}); err != nil {
		return unit, err
	}

	return pgmodel.ParseLinkUnitModels(unitModel, linkUnitInfoModel)
}

func (repo *linkUnitRepo) Update(linkUnit linkunitmodel.LinkUnit) error {
	linkUnitInfoModel := pgmodel.NewLinkUnitInfoModel(linkUnit)
	unitModel := pgmodel.NewLinkUnitModel(linkUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(linkUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			linkUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}
		if err := transaction.Save(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Save(&linkUnitInfoModel).Error
	})
}

func (repo *linkUnitRepo) Delete(linkUnit linkunitmodel.LinkUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			linkUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Where(
			"id = ?",
			linkUnit.GetId().Uuid(),
		).Delete(&pgmodel.UnitModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			linkUnit.GetId().Uuid(),
		).Delete(&pgmodel.LinkUnitInfoModel{}).Error
	})
}
