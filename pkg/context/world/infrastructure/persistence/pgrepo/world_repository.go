package pgrepo

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/worldmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/commonutil"
	"gorm.io/gorm"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/infrastructure/persistence/pgmodel"
)

func newModelFromWorld(world worldmodel.World) pgmodel.WorldModel {
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

func parseModelToWorld(worldModel pgmodel.WorldModel) (world worldmodel.World, err error) {
	bound, err := worldcommonmodel.NewBound(
		worldcommonmodel.NewPosition(worldModel.BoundFromX, worldModel.BoundFromZ),
		worldcommonmodel.NewPosition(worldModel.BoundToX, worldModel.BoundToZ),
	)
	if err != nil {
		return world, err
	}
	return worldmodel.LoadWorld(
		globalcommonmodel.NewWorldId(worldModel.Id),
		globalcommonmodel.NewUserId(worldModel.UserId),
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
	worldModel := newModelFromWorld(world)
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Create(&worldModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&world)
}

func (repo *worldRepo) Update(world worldmodel.World) error {
	worldModel := newModelFromWorld(world)
	worldModel.UpdatedAt = time.Now()
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Model(&pgmodel.WorldModel{}).Where(
			"id = ?",
			world.GetId().Uuid(),
		).Select("*").Updates(&worldModel).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&world)
}

func (repo *worldRepo) Delete(world worldmodel.World) error {
	if err := repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Delete(&pgmodel.WorldModel{}, world.GetId().Uuid()).Error
	}); err != nil {
		return err
	}
	return repo.domainEventDispatcher.Dispatch(&world)
}

func (repo *worldRepo) Get(worldId globalcommonmodel.WorldId) (world worldmodel.World, err error) {
	worldModel := pgmodel.WorldModel{Id: worldId.Uuid()}
	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.First(&worldModel).Error
	}); err != nil {
		return world, err
	}
	return parseModelToWorld(worldModel)
}

func (repo *worldRepo) Query(limit int, offset int) (worlds []worldmodel.World, err error) {
	var worldModels []pgmodel.WorldModel

	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Limit(limit).Offset(offset).Order("created_at desc").Find(&worldModels).Error
	}); err != nil {
		return worlds, err
	}

	return commonutil.MapWithError[pgmodel.WorldModel](worldModels, func(_ int, worldModel pgmodel.WorldModel) (worldmodel.World, error) {
		return parseModelToWorld(worldModel)
	})
}

func (repo *worldRepo) GetWorldsOfUser(userId globalcommonmodel.UserId) (worlds []worldmodel.World, err error) {
	var worldModels []pgmodel.WorldModel

	if err = repo.uow.Execute(func(transaction *gorm.DB) error {
		return transaction.Order("created_at desc").Find(
			&worldModels,
			pgmodel.WorldModel{
				UserId: userId.Uuid(),
			},
		).Error
	}); err != nil {
		return worlds, err
	}

	return commonutil.MapWithError[pgmodel.WorldModel](worldModels, func(_ int, worldModel pgmodel.WorldModel) (worldmodel.World, error) {
		return parseModelToWorld(worldModel)
	})
}
