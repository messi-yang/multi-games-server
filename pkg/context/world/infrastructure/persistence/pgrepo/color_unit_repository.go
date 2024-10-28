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
	colorUnitInfoModel := pgmodel.NewColorUnitInfoModel(colorUnit)
	unitModel := pgmodel.NewColorUnitModel(colorUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(colorUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&colorUnitInfoModel).Error; err != nil {
			return err
		}
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *colorUnitRepo) Get(id colorunitmodel.ColorUnitId) (unit colorunitmodel.ColorUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	colorUnitInfoModel := pgmodel.ColorUnitInfoModel{}

	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumColor,
		).First(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&colorUnitInfoModel).Error
	}); err != nil {
		return unit, err
	}

	return pgmodel.ParseColorUnitModels(unitModel, colorUnitInfoModel)
}

func (repo *colorUnitRepo) Update(colorUnit colorunitmodel.ColorUnit) error {
	colorUnitInfoModel := pgmodel.NewColorUnitInfoModel(colorUnit)
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
		if err := transaction.Model(&pgmodel.UnitModel{}).Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			colorUnit.GetWorldId().Uuid(),
			colorUnit.GetPosition().GetX(),
			colorUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumColor,
		).Select("*").Updates(unitModel).Error; err != nil {
			return err
		}
		return transaction.Model(&pgmodel.ColorUnitInfoModel{}).Where(
			"id = ?",
			colorUnit.GetId().Uuid(),
		).Select("*").Updates(colorUnitInfoModel).Error
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
		if err := transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ? AND type = ?",
			colorUnit.GetWorldId().Uuid(),
			colorUnit.GetPosition().GetX(),
			colorUnit.GetPosition().GetZ(),
			pgmodel.UnitTypeEnumColor,
		).Delete(&pgmodel.UnitModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			colorUnit.GetId().Uuid(),
		).Delete(&pgmodel.ColorUnitInfoModel{}).Error
	})
}
