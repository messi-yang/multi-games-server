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

func NewUnitRepository() (unitmodel.Repository, error) {
	gormDb, err := NewSession()
	if err != nil {
		return nil, err
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

func (repo *unitRepository) GetUnitAt(worldId worldmodel.WorldIdVo, position commonmodel.PositionVo) (unitmodel.UnitAgg, bool, error) {
	unitModel := psqlmodel.UnitModel{}
	result := repo.gormDb.Where(
		"world_id = ? AND pos_x = ? AND pos_z = ?",
		worldId.Uuid(),
		position.GetX(),
		position.GetZ(),
	).First(&unitModel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return unitmodel.UnitAgg{}, false, nil
		}
		return unitmodel.UnitAgg{}, false, result.Error
	}

	return unitModel.ToAggregate(), true, nil
}

func (repo *unitRepository) GetUnitsInBound(worldId worldmodel.WorldIdVo, bound commonmodel.BoundVo) ([]unitmodel.UnitAgg, error) {
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
		return nil, result.Error
	}
	return lo.Map(unitModels, func(model psqlmodel.UnitModel, _ int) unitmodel.UnitAgg {
		return model.ToAggregate()
	}), nil
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
