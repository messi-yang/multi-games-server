package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/worldmodel"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/samber/lo"
)

func newWorldModel(world worldmodel.World) pgmodel.WorldModel {
	return pgmodel.WorldModel{
		Id:     world.GetId().Uuid(),
		UserId: world.GetUserId().Uuid(),
		Name:   world.GetName(),
	}
}

func parseWorldModel(worldModel pgmodel.WorldModel) worldmodel.World {
	return worldmodel.NewWorld(commonmodel.NewWorldId(worldModel.Id), sharedkernelmodel.NewUserId(worldModel.UserId), worldModel.Name)
}

type worldRepo struct {
	uow pguow.Uow
}

func NewWorldRepo(uow pguow.Uow) (repository worldmodel.Repo) {
	return &worldRepo{uow: uow}
}

func (repo *worldRepo) Add(world worldmodel.World) error {
	worldModel := newWorldModel(world)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldModel).Error
	})
}

func (repo *worldRepo) Update(world worldmodel.World) error {
	worldModel := newWorldModel(world)
	return repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&worldModel).Error
	})
}

func (repo *worldRepo) Get(worldId commonmodel.WorldId) (world worldmodel.World, err error) {
	worldModel := pgmodel.WorldModel{Id: worldId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldModel).Error
	}); err != nil {
		return world, err
	}
	return parseWorldModel(worldModel), nil
}

func (repo *worldRepo) GetAll() (worlds []worldmodel.World, err error) {
	var worldModels []pgmodel.WorldModel
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(&worldModels).Limit(10).Error
	}); err != nil {
		return worlds, err
	}

	worlds = lo.Map(worldModels, func(worldModel pgmodel.WorldModel, _ int) worldmodel.World {
		return parseWorldModel(worldModel)
	})
	return worlds, nil
}
