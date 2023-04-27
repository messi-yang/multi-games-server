package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newUnitModel(unit unitmodel.UnitAgg) pgmodel.UnitModel {
	return pgmodel.UnitModel{
		WorldId:   unit.GetWorldId().Uuid(),
		PosX:      unit.GetPosition().GetX(),
		PosZ:      unit.GetPosition().GetZ(),
		ItemId:    unit.GetItemId().Uuid(),
		Direction: unit.GetDirection().Int8(),
	}
}

func parseUnitModel(unitModel pgmodel.UnitModel) unitmodel.UnitAgg {
	return unitmodel.NewUnitAgg(
		commonmodel.NewWorldIdVo(unitModel.WorldId),
		commonmodel.NewPositionVo(unitModel.PosX, unitModel.PosZ),
		commonmodel.NewItemIdVo(unitModel.ItemId),
		commonmodel.NewDirectionVo(unitModel.Direction),
	)
}

type unitRepo struct {
	db *gorm.DB
}

func NewUnitRepo() (repository unitmodel.Repo, err error) {
	db, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &unitRepo{db: db}, nil
}

func (repo *unitRepo) Add(unit unitmodel.UnitAgg) error {
	unitModel := newUnitModel(unit)
	res := repo.db.Create(&unitModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *unitRepo) GetUnitAt(
	worldId commonmodel.WorldIdVo, position commonmodel.PositionVo,
) (unit unitmodel.UnitAgg, found bool, err error) {
	unitModels := []pgmodel.UnitModel{}
	result := repo.db.Where(
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

func (repo *unitRepo) GetUnitsInBound(
	worldId commonmodel.WorldIdVo, bound commonmodel.BoundVo,
) (units []unitmodel.UnitAgg, err error) {
	var unitModels []pgmodel.UnitModel
	result := repo.db.Where(
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
	units = lo.Map(unitModels, func(unitModel pgmodel.UnitModel, _ int) unitmodel.UnitAgg {
		return parseUnitModel(unitModel)
	})
	return units, nil
}

func (repo *unitRepo) Delete(worldId commonmodel.WorldIdVo, position commonmodel.PositionVo) error {
	result := repo.db.Where(
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
