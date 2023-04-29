package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/application/uow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
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
	uow uow.Uow[*gorm.DB]
}

func NewUnitRepo(uow uow.Uow[*gorm.DB]) (repository unitmodel.Repo) {
	return &unitRepo{uow: uow}
}

func (repo *unitRepo) Add(unit unitmodel.Unit) error {
	unitModel := newUnitModel(unit)
	res := repo.uow.GetTransaction().Create(&unitModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *unitRepo) FindUnitAt(
	worldId commonmodel.WorldId, position commonmodel.Position,
) (unit unitmodel.Unit, found bool, err error) {
	unitModels := []pgmodel.UnitModel{}
	result := repo.uow.GetTransaction().Where(
		"world_id = ? AND pos_x = ? AND pos_z = ?",
		worldId.Uuid(),
		position.GetX(),
		position.GetZ(),
	).Find(&unitModels)

	if result.Error != nil {
		return unit, found, result.Error
	}

	found = result.RowsAffected >= 1
	if found {
		unit = parseUnitModel(unitModels[0])
	}
	return unit, found, nil
}

func (repo *unitRepo) QueryUnitsInBound(
	worldId commonmodel.WorldId, bound commonmodel.Bound,
) (units []unitmodel.Unit, err error) {
	var unitModels []pgmodel.UnitModel
	result := repo.uow.GetTransaction().Where(
		"world_id = ? AND pos_x >= ? AND pos_x <= ? AND pos_z >= ? AND pos_z <= ?",
		worldId.Uuid(),
		bound.GetFrom().GetX(),
		bound.GetTo().GetX(),
		bound.GetFrom().GetZ(),
		bound.GetTo().GetZ(),
	).Find(&unitModels, pgmodel.UnitModel{})
	if result.Error != nil {
		return units, result.Error
	}
	units = lo.Map(unitModels, func(unitModel pgmodel.UnitModel, _ int) unitmodel.Unit {
		return parseUnitModel(unitModel)
	})
	return units, nil
}

func (repo *unitRepo) Delete(worldId commonmodel.WorldId, position commonmodel.Position) error {
	result := repo.uow.GetTransaction().Where(
		"world_id = ? AND pos_x = ? AND pos_z = ?",
		worldId.Uuid(),
		position.GetX(),
		position.GetZ(),
	).Delete(&pgmodel.UnitModel{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
