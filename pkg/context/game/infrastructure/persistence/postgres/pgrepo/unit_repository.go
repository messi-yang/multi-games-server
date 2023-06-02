package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel/unitmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/samber/lo"
)

func newUnitModel(unit unitmodel.Unit) pgmodel.UnitModel {
	return pgmodel.UnitModel{
		WorldId:   unit.GetWorldId().Uuid(),
		PosX:      unit.GetPosition().GetX(),
		PosZ:      unit.GetPosition().GetZ(),
		ItemId:    unit.GetItemId().Uuid(),
		Direction: unit.GetDirection().Int8(),
	}
}

func parseUnitModel(unitModel pgmodel.UnitModel) unitmodel.Unit {
	worldId := sharedkernelmodel.NewWorldId(unitModel.WorldId)
	pos := commonmodel.NewPosition(unitModel.PosX, unitModel.PosZ)
	return unitmodel.LoadUnit(
		unitmodel.NewUnitId(worldId, pos),
		worldId,
		pos,
		commonmodel.NewItemId(unitModel.ItemId),
		commonmodel.NewDirection(unitModel.Direction),
	)
}

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

func (repo *unitRepo) Add(unit unitmodel.Unit) error {
	unitModel := newUnitModel(unit)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&unitModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&unit)
}

func (repo *unitRepo) Delete(unit unitmodel.Unit) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			unit.GetWorldId().Uuid(),
			unit.GetPosition().GetX(),
			unit.GetPosition().GetZ(),
		).Delete(&pgmodel.UnitModel{}).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&unit)
}

func (repo *unitRepo) FindUnitAt(
	worldId sharedkernelmodel.WorldId, position commonmodel.Position,
) (unit unitmodel.Unit, found bool, err error) {
	unitModels := []pgmodel.UnitModel{}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			worldId.Uuid(),
			position.GetX(),
			position.GetZ(),
		).Find(&unitModels).Error
	}); err != nil {
		return unit, found, err
	}

	found = len(unitModels) >= 1
	if found {
		unit = parseUnitModel(unitModels[0])
	}
	return unit, found, nil
}

func (repo *unitRepo) QueryUnitsInBound(
	worldId sharedkernelmodel.WorldId, bound commonmodel.Bound,
) (units []unitmodel.Unit, err error) {
	var unitModels []pgmodel.UnitModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x >= ? AND pos_x <= ? AND pos_z >= ? AND pos_z <= ?",
			worldId.Uuid(),
			bound.GetFrom().GetX(),
			bound.GetTo().GetX(),
			bound.GetFrom().GetZ(),
			bound.GetTo().GetZ(),
		).Find(&unitModels, pgmodel.UnitModel{}).Error
	}); err != nil {
		return units, err
	}

	units = lo.Map(unitModels, func(unitModel pgmodel.UnitModel, _ int) unitmodel.Unit {
		return parseUnitModel(unitModel)
	})
	return units, nil
}
