package pgrepo

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel/portalunitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type portalUnitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewPortalUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository portalunitmodel.PortalUnitRepo) {
	return &portalUnitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *portalUnitRepo) Add(portalUnit portalunitmodel.PortalUnit) error {
	portalUnitInfoModel := pgmodel.NewPortalUnitInfoModel(portalUnit)
	unitModel := pgmodel.NewPortalUnitModel(portalUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(portalUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Create(&portalUnitInfoModel).Error; err != nil {
			return err
		}
		if err := transaction.Create(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Create(occupiedPositionModels).Error
	})
}

func (repo *portalUnitRepo) Get(id portalunitmodel.PortalUnitId) (unit portalunitmodel.PortalUnit, err error) {
	unitModel := pgmodel.UnitModel{}
	portalUnitInfoModel := pgmodel.PortalUnitInfoModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"id = ? AND type = ?",
			id.Uuid(),
			pgmodel.UnitTypeEnumPortal,
		).First(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			unitModel.Id,
		).First(&portalUnitInfoModel).Error
	}); err != nil {
		return unit, err
	}

	return pgmodel.ParsePortalUnitModels(unitModel, portalUnitInfoModel)
}

func (repo *portalUnitRepo) Update(portalUnit portalunitmodel.PortalUnit) error {
	portalUnitInfoModel := pgmodel.NewPortalUnitInfoModel(portalUnit)
	unitModel := pgmodel.NewPortalUnitModel(portalUnit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(portalUnit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			portalUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}
		if err := transaction.Save(&unitModel).Error; err != nil {
			return err
		}
		return transaction.Save(&portalUnitInfoModel).Error
	})
}

func (repo *portalUnitRepo) Query(worldId globalcommonmodel.WorldId, limit int, offset int) ([]portalunitmodel.PortalUnit, error) {
	// query units first
	var unitModels []pgmodel.UnitModel
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND type = ?",
			worldId.Uuid(),
			pgmodel.UnitTypeEnumPortal,
		).Order("pos_z asc, pos_x asc").Limit(limit).Offset(offset).Find(&unitModels).Error
	}); err != nil {
		return nil, err
	}

	unitModelIds := make([]uuid.UUID, len(unitModels))
	for i, unitModel := range unitModels {
		unitModelIds[i] = unitModel.Id
	}

	var portalUnitInfoModels []pgmodel.PortalUnitInfoModel
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id IN (?)",
			unitModelIds,
		).Find(&portalUnitInfoModels).Error
	}); err != nil {
		return nil, err
	}

	fmt.Println("unitModelIds", unitModelIds)
	fmt.Println("portalUnitInfoModels", portalUnitInfoModels)

	portalUnits := make([]portalunitmodel.PortalUnit, len(unitModels))
	for i, unitModel := range unitModels {
		var err error
		portalUnits[i], err = pgmodel.ParsePortalUnitModels(unitModel, portalUnitInfoModels[i])
		if err != nil {
			return nil, err
		}
	}

	return portalUnits, nil
}

func (repo *portalUnitRepo) Delete(portalUnit portalunitmodel.PortalUnit) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			portalUnit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Where(
			"id = ?",
			portalUnit.GetId().Uuid(),
		).Delete(&pgmodel.UnitModel{}).Error; err != nil {
			return err
		}
		return transaction.Where(
			"id = ?",
			portalUnit.GetId().Uuid(),
		).Delete(&pgmodel.PortalUnitInfoModel{}).Error
	})
}

func (repo *portalUnitRepo) GetTopLeftMostUnitWithoutTarget(worldId globalcommonmodel.WorldId) (portalUnit *portalunitmodel.PortalUnit, err error) {
	var portalUnitInfoModels []pgmodel.PortalUnitInfoModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Select("units.pos_x, units.pos_z, portal_unit_infos.*").Joins(
			"left join units on units.id = portal_unit_infos.id",
		).Where(
			"portal_unit_infos.world_id = ? AND portal_unit_infos.target_unit_id IS NULL",
			worldId.Uuid(),
		).Order("units.pos_z asc, units.pos_x asc").Limit(1).Find(&portalUnitInfoModels).Error
	}); err != nil {
		return portalUnit, err
	}

	if len(portalUnitInfoModels) == 0 {
		return nil, nil
	}

	var unitModel pgmodel.UnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ?",
			portalUnitInfoModels[0].Id,
		).First(&unitModel).Error
	}); err != nil {
		return portalUnit, err
	}

	firstPortalUnitWithNoTarget, err := pgmodel.ParsePortalUnitModels(unitModel, portalUnitInfoModels[0])
	if err != nil {
		return nil, err
	}
	return commonutil.ToPointer(firstPortalUnitWithNoTarget), err
}
