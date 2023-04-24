package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newWorldModel(world worldmodel.WorldAgg) pgmodel.WorldModel {
	return pgmodel.WorldModel{
		Id:      world.GetId().Uuid(),
		GamerId: world.GetGamerId().Uuid(),
		Name:    world.GetName(),
	}
}

func parseWorldModel(worldModel pgmodel.WorldModel) worldmodel.WorldAgg {
	return worldmodel.NewWorldAgg(commonmodel.NewWorldIdVo(worldModel.Id), commonmodel.NewGamerIdVo(worldModel.GamerId))
}

type worldRepo struct {
	dbClient *gorm.DB
}

func NewWorldRepo() (repository worldmodel.Repo, err error) {
	dbClient, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &worldRepo{dbClient: dbClient}, nil
}

func (repo *worldRepo) Get(worldId commonmodel.WorldIdVo) (world worldmodel.WorldAgg, err error) {
	worldModel := pgmodel.WorldModel{Id: worldId.Uuid()}
	result := repo.dbClient.First(&worldModel)
	if result.Error != nil {
		return world, result.Error
	}
	return parseWorldModel(worldModel), nil
}

func (repo *worldRepo) GetAll() (worlds []worldmodel.WorldAgg, err error) {
	var worldModels []pgmodel.WorldModel
	result := repo.dbClient.Find(&worldModels).Limit(10)
	if result.Error != nil {
		return worlds, result.Error
	}

	worlds = lo.Map(worldModels, func(worldModel pgmodel.WorldModel, _ int) worldmodel.WorldAgg {
		return parseWorldModel(worldModel)
	})
	return worlds, nil
}

func (repo *worldRepo) Add(world worldmodel.WorldAgg) error {
	worldModel := newWorldModel(world)
	res := repo.dbClient.Create(&worldModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *worldRepo) ReadLockAccess(worldId commonmodel.WorldIdVo) func() {
	return func() {}
}

func (repo *worldRepo) LockAccess(worldId commonmodel.WorldIdVo) func() {
	return func() {}
}
