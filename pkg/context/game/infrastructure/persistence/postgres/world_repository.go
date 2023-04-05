package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/infrastructure/persistence/postgres/psqlmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type worldRepository struct {
	dbClient *gorm.DB
}

func NewWorldRepository() (repository worldmodel.Repository, err error) {
	dbClient, err := NewDbClient()
	if err != nil {
		return repository, err
	}
	return &worldRepository{dbClient: dbClient}, nil
}

func (repo *worldRepository) Get(worldId worldmodel.WorldIdVo) (world worldmodel.WorldAgg, err error) {
	worldModel := psqlmodel.WorldModel{Id: worldId.Uuid()}
	result := repo.dbClient.First(&worldModel)
	if result.Error != nil {
		return world, result.Error
	}
	return worldModel.ToAggregate(), nil
}

func (repo *worldRepository) GetWorldOfUser(userId usermodel.UserIdVo) (world worldmodel.WorldAgg, found bool, err error) {
	worldModels := []psqlmodel.WorldModel{}
	result := repo.dbClient.Where("user_id = ?", userId.Uuid()).Find(&worldModels)
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
	result := repo.dbClient.Find(&worldModels).Limit(10)
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
	res := repo.dbClient.Create(&worldModel)
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
