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

func NewWorldRepository() (repository worldmodel.Repository, err error) {
	gormDb, err := NewSession()
	if err != nil {
		return repository, err
	}
	return &worldRepository{gormDb: gormDb}, nil
}

func (repo *worldRepository) Get(worldId worldmodel.WorldIdVo) (world worldmodel.WorldAgg, err error) {
	worldModel := psqlmodel.WorldModel{Id: worldId.Uuid()}
	result := repo.gormDb.First(&worldModel)
	if result.Error != nil {
		return world, result.Error
	}
	return worldModel.ToAggregate(), nil
}

func (repo *worldRepository) GetWorldOfUser(userId usermodel.UserIdVo) (world worldmodel.WorldAgg, found bool, err error) {
	worldModels := []psqlmodel.WorldModel{}
	result := repo.gormDb.Where("user_id = ?", userId.Uuid()).Find(&worldModels)
	if result.Error != nil {
		return world, found, result.Error
	}
	found = result.RowsAffected > 0
	if found {
		world = worldModels[0].ToAggregate()
	}
	return world, found, nil
}

func (repo *worldRepository) GetAll() (worlds []worldmodel.WorldAgg, err error) {
	var worldModels []psqlmodel.WorldModel
	result := repo.gormDb.Find(&worldModels).Limit(10)
	if result.Error != nil {
		return worlds, result.Error
	}

	worlds = lo.Map(worldModels, func(model psqlmodel.WorldModel, _ int) worldmodel.WorldAgg {
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
