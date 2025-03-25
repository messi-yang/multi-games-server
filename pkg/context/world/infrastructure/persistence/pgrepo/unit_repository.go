package pgrepo

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/blockmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/unitmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

type unitRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewUnitRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository unitmodel.UnitRepo) {
	return &unitRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *unitRepo) Get(id unitmodel.UnitId) (unit unitmodel.Unit, err error) {
	unitModel := pgmodel.UnitModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"id = ?",
			id.Uuid(),
		).First(&unitModel).Error
	}); err != nil {
		return unit, err
	}
	return pgmodel.ParseUnitModel(unitModel)
}

func (repo *unitRepo) HasUnitsInBound(worldId globalcommonmodel.WorldId, bound worldcommonmodel.Bound) (bool, error) {
	occupiedPositionModels := []pgmodel.OccupiedPositionModel{}
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x >= ? AND pos_z >= ? AND pos_x <= ? AND pos_z <= ?",
			worldId.Uuid(),
			bound.GetFrom().GetX(),
			bound.GetFrom().GetZ(),
			bound.GetTo().GetX(),
			bound.GetTo().GetZ(),
		).Find(&occupiedPositionModels, pgmodel.OccupiedPositionModel{}).Error
	}); err != nil {
		return false, err
	}

	return len(occupiedPositionModels) > 0, nil
}

func (repo *unitRepo) Update(unit unitmodel.Unit) error {
	unitModel := pgmodel.NewUnitModel(unit)
	occupiedPositionModels := pgmodel.NewOccupiedPositionModels(unit.UnitEntity)

	return repo.uow.Execute(func(transaction *gorm.DB) error {
		if err := transaction.Where(
			"unit_id = ?",
			unit.GetId().Uuid(),
		).Delete(&pgmodel.OccupiedPositionModel{}).Error; err != nil {
			return err
		}
		if err := transaction.Create(occupiedPositionModels).Error; err != nil {
			return err
		}

		return transaction.Save(&unitModel).Error
	})
}

func (repo *unitRepo) GetUnitsOfWorld(
	worldId globalcommonmodel.WorldId,
) (units []unitmodel.Unit, err error) {
	var unitModels []pgmodel.UnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ?",
			worldId.Uuid(),
		).Find(&unitModels, pgmodel.UnitModel{}).Error
	}); err != nil {
		return units, err
	}

	return commonutil.MapWithError(unitModels, func(_ int, unitModel pgmodel.UnitModel) (unitmodel.Unit, error) {
		return pgmodel.ParseUnitModel(unitModel)
	})
}

func (repo *unitRepo) GetUnitsInBlock(worldId globalcommonmodel.WorldId, block blockmodel.Block) (units []unitmodel.Unit, err error) {
	blockFrom := block.GetBound().GetFrom()
	blockTo := block.GetBound().GetTo()

	var unitModels []pgmodel.UnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x >= ? AND pos_z >= ? AND pos_x <= ? AND pos_z <= ?",
			worldId.Uuid(),
			blockFrom.GetX(),
			blockFrom.GetZ(),
			blockTo.GetX(),
			blockTo.GetZ(),
		).Find(&unitModels, pgmodel.UnitModel{}).Error
	}); err != nil {
		return units, err
	}

	return commonutil.MapWithError(unitModels, func(_ int, unitModel pgmodel.UnitModel) (unitmodel.Unit, error) {
		return pgmodel.ParseUnitModel(unitModel)
	})
}
