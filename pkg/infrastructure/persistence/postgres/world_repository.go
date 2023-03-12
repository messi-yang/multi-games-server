package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type worldRepository struct {
	gormDb *gorm.DB
}

func NewWorldRepository() (worldmodel.Repository, error) {
	gormDb, err := NewSession()
	if err != nil {
		return nil, err
	}
	return &worldRepository{gormDb: gormDb}, nil
}

func (repo *worldRepository) ExistsWithUserId(userId usermodel.UserIdVo) (bool, error) {
	worldModel := psqlmodel.WorldModel{UserId: userId.Uuid()}
	result := repo.gormDb.First(&worldModel)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func (repo *worldRepository) GetAll() ([]worldmodel.WorldAgg, error) {
	var worldModels []psqlmodel.WorldModel
	result := repo.gormDb.Select("Id", "Width", "Height", "CreatedAt", "UpdatedAt").Find(&worldModels)
	if result.Error != nil {
		return nil, result.Error
	}

	worlds := lo.Map(worldModels, func(model psqlmodel.WorldModel, _ int) worldmodel.WorldAgg {
		return model.ToAggregate()
	})

	return worlds, nil
}

func (repo *worldRepository) Add(world worldmodel.WorldAgg) error {
	worldModel := psqlmodel.NewWorldModel(world)
	res := repo.gormDb.Create(&worldModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *worldRepository) ReadLockAccess(worldId worldmodel.WorldIdVo) func() {
	return func() {}
}

func (repo *worldRepository) LockAccess(worldId worldmodel.WorldIdVo) func() {
	return func() {}
}
