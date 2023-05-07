package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"gorm.io/gorm"

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
	return unitmodel.NewUnit(
		commonmodel.NewWorldId(unitModel.WorldId),
		commonmodel.NewPosition(unitModel.PosX, unitModel.PosZ),
		commonmodel.NewItemId(unitModel.ItemId),
		commonmodel.NewDirection(unitModel.Direction),
	)
}

type unitRepo struct {
	uow pguow.Uow
}

func NewUnitRepo(uow pguow.Uow) (repository unitmodel.Repo) {
	return &unitRepo{uow: uow}
}

func (repo *unitRepo) Add(unit unitmodel.Unit) error {
	unitModel := newUnitModel(unit)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&unitModel).Error
	})
}

func (repo *unitRepo) Delete(worldId commonmodel.WorldId, position commonmodel.Position) error {
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Where(
			"world_id = ? AND pos_x = ? AND pos_z = ?",
			worldId.Uuid(),
			position.GetX(),
			position.GetZ(),
		).Delete(&pgmodel.UnitModel{}).Error
	})
}

func (repo *unitRepo) FindUnitAt(
	worldId commonmodel.WorldId, position commonmodel.Position,
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
	worldId commonmodel.WorldId, bound commonmodel.Bound,
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
