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
		return
	}
	repository = &worldRepository{gormDb: gormDb}
	return
}

func (repo *worldRepository) Get(worldId worldmodel.WorldIdVo) (world worldmodel.WorldAgg, err error) {
	worldModel := psqlmodel.WorldModel{Id: worldId.Uuid()}
	result := repo.gormDb.First(&worldModel)
	if result.Error != nil {
		err = result.Error
	}
	world = worldModel.ToAggregate()
	return
}

func (repo *worldRepository) GetWorldOfUser(userId usermodel.UserIdVo) (world worldmodel.WorldAgg, found bool, err error) {
	worldModels := []psqlmodel.WorldModel{}
	result := repo.gormDb.Where("user_id = ?", userId.Uuid()).Find(&worldModels)
	if result.Error != nil {
		err = result.Error
		return
	}
	found = result.RowsAffected > 0
	if found {
		world = worldModels[0].ToAggregate()
	}
	return
}

func (repo *worldRepository) GetAll() (worlds []worldmodel.WorldAgg, err error) {
	var worldModels []psqlmodel.WorldModel
	result := repo.gormDb.Select("Id", "Width", "Height", "CreatedAt", "UpdatedAt").Find(&worldModels)
	if result.Error != nil {
		err = result.Error
		return
	}

	worlds = lo.Map(worldModels, func(model psqlmodel.WorldModel, _ int) worldmodel.WorldAgg {
		return model.ToAggregate()
	})
	return
}

func (repo *worldRepository) Add(world worldmodel.WorldAgg) (err error) {
	worldModel := psqlmodel.NewWorldModel(world)
	res := repo.gormDb.Create(&worldModel)
	if res.Error != nil {
		err = res.Error
		return
	}
	return
}

func (repo *worldRepository) ReadLockAccess(worldId worldmodel.WorldIdVo) func() {
	return func() {}
}

func (repo *worldRepository) LockAccess(worldId worldmodel.WorldIdVo) func() {
	return func() {}
}
