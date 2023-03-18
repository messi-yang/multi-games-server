package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type unitRepository struct {
	gormDb *gorm.DB
}

func NewUnitRepository() (repository unitmodel.Repository, err error) {
	gormDb, err := NewSession()
	if err != nil {
		return repository, err
	}
	return &unitRepository{gormDb: gormDb}, nil
}

func (repo *unitRepository) Add(unit unitmodel.UnitAgg) error {
	unitModel := psqlmodel.NewUnitModel(unit)
	res := repo.gormDb.Create(&unitModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *unitRepository) GetUnitAt(
	worldId worldmodel.WorldIdVo, position commonmodel.PositionVo,
) (unit unitmodel.UnitAgg, found bool, err error) {
	unitModels := []psqlmodel.UnitModel{}
	result := repo.gormDb.Where(
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
		unit = unitModels[0].ToAggregate()
	}
	return unit, found, nil
}

func (repo *unitRepository) GetUnitsInBound(
	worldId worldmodel.WorldIdVo, bound commonmodel.BoundVo,
) (units []unitmodel.UnitAgg, err error) {
	var unitModels []psqlmodel.UnitModel
	result := repo.gormDb.Where(
		"world_id = ? AND pos_x >= ? AND pos_x <= ? AND pos_z >= ? AND pos_z <= ?",
		worldId.Uuid(),
		bound.GetFrom().GetX(),
		bound.GetTo().GetX(),
		bound.GetFrom().GetZ(),
		bound.GetTo().GetZ(),
	).Find(&unitModels, psqlmodel.UnitModel{})
	if result.Error != nil {
		return units, result.Error
	}
	units = lo.Map(unitModels, func(model psqlmodel.UnitModel, _ int) unitmodel.UnitAgg {
		return model.ToAggregate()
	})
	return units, nil
}

func (repo *unitRepository) Delete(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) error {
	result := repo.gormDb.Where(
		"world_id = ? AND pos_x = ? AND pos_z = ?",
		worldId.Uuid(),
		position.GetX(),
		position.GetZ(),
	).Delete(&psqlmodel.UnitModel{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
