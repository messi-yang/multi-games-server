package pgrepo

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pgmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
)

func newWorldModel(world worldmodel.World) pgmodel.WorldModel {
	return pgmodel.WorldModel{
		Id:         world.GetId().Uuid(),
		UserId:     world.GetUserId().Uuid(),
		Name:       world.GetName(),
		BoundFromX: world.GetBound().GetFrom().GetX(),
		BoundFromZ: world.GetBound().GetFrom().GetZ(),
		BoundToX:   world.GetBound().GetTo().GetX(),
		BoundToZ:   world.GetBound().GetTo().GetZ(),
		UpdatedAt:  world.GetUpdatedAt(),
		CreatedAt:  world.GetCreatedAt(),
	}
}

func parseWorldModel(worldModel pgmodel.WorldModel) (world worldmodel.World, err error) {
	bound, err := commonmodel.NewBound(
		commonmodel.NewPosition(worldModel.BoundFromX, worldModel.BoundFromZ),
		commonmodel.NewPosition(worldModel.BoundToX, worldModel.BoundToZ),
	)
	if err != nil {
		return world, err
	}
	return worldmodel.LoadWorld(
		sharedkernelmodel.NewWorldId(worldModel.Id),
		sharedkernelmodel.NewUserId(worldModel.UserId),
		worldModel.Name,
		bound,
		worldModel.CreatedAt,
		worldModel.UpdatedAt,
	), nil
}

type worldRepo struct {
	uow                   pguow.Uow
	domainEventDispatcher domain.DomainEventDispatcher
}

func NewWorldRepo(uow pguow.Uow, domainEventDispatcher domain.DomainEventDispatcher) (repository worldmodel.WorldRepo) {
	return &worldRepo{
		uow:                   uow,
		domainEventDispatcher: domainEventDispatcher,
	}
}

func (repo *worldRepo) Add(world worldmodel.World) error {
	worldModel := newWorldModel(world)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&world)
}

func (repo *worldRepo) Update(world worldmodel.World) error {
	worldModel := newWorldModel(world)
	worldModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Save(&worldModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&world)
}

func (repo *worldRepo) Get(worldId sharedkernelmodel.WorldId) (world worldmodel.World, err error) {
	worldModel := pgmodel.WorldModel{Id: worldId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldModel).Error
	}); err != nil {
		return world, err
	}
	return parseWorldModel(worldModel)
}

func (repo *worldRepo) Query(limit int, offset int) (worlds []worldmodel.World, err error) {
	var worldModels []pgmodel.WorldModel

	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Limit(limit).Offset(offset).Order("created_at DESC").Find(&worldModels).Error
	}); err != nil {
		return worlds, err
	}

	return commonutil.MapWithError[pgmodel.WorldModel](worldModels, func(_ int, worldModel pgmodel.WorldModel) (worldmodel.World, error) {
		return parseWorldModel(worldModel)
	})
}

func (repo *worldRepo) GetWorldsOfUser(userId sharedkernelmodel.UserId) (worlds []worldmodel.World, err error) {
	var worldModels []pgmodel.WorldModel

	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Find(
			&worldModels,
			pgmodel.WorldModel{
				UserId: userId.Uuid(),
			},
		).Error
	}); err != nil {
		return worlds, err
	}

	return commonutil.MapWithError[pgmodel.WorldModel](worldModels, func(_ int, worldModel pgmodel.WorldModel) (worldmodel.World, error) {
		return parseWorldModel(worldModel)
	})
}
