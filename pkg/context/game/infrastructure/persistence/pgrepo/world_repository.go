package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pguow"
	"github.com/samber/lo"
)

func newWorldModel(world worldmodel.World) pgmodel.WorldModel {
	return pgmodel.WorldModel{
		Id:      world.GetId().Uuid(),
		GamerId: world.GetGamerId().Uuid(),
		Name:    world.GetName(),
	}
}

func parseWorldModel(worldModel pgmodel.WorldModel) worldmodel.World {
	return worldmodel.NewWorld(commonmodel.NewWorldId(worldModel.Id), commonmodel.NewGamerId(worldModel.GamerId), worldModel.Name)
}

type worldRepo struct {
	uow pguow.Uow
}

func NewWorldRepo(uow pguow.Uow) (repository worldmodel.Repo) {
	return &worldRepo{uow: uow}
}

func (repo *worldRepo) Add(world worldmodel.World) error {
	worldModel := newWorldModel(world)
	res := repo.uow.GetTransaction().Create(&worldModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *worldRepo) Update(world worldmodel.World) error {
	worldModel := newWorldModel(world)
	res := repo.uow.GetTransaction().Save(&worldModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *worldRepo) Get(worldId commonmodel.WorldId) (world worldmodel.World, err error) {
	worldModel := pgmodel.WorldModel{Id: worldId.Uuid()}
	result := repo.uow.GetTransaction().First(&worldModel)
	if result.Error != nil {
		return world, result.Error
	}
	return parseWorldModel(worldModel), nil
}

func (repo *worldRepo) GetAll() (worlds []worldmodel.World, err error) {
	var worldModels []pgmodel.WorldModel
	result := repo.uow.GetTransaction().Find(&worldModels).Limit(10)
	if result.Error != nil {
		return worlds, result.Error
	}

	worlds = lo.Map(worldModels, func(worldModel pgmodel.WorldModel, _ int) worldmodel.World {
		return parseWorldModel(worldModel)
	})
	return worlds, nil
}

func (repo *worldRepo) ReadLockAccess(worldId commonmodel.WorldId) func() {
	return func() {}
}

func (repo *worldRepo) LockAccess(worldId commonmodel.WorldId) func() {
	return func() {}
}
