package postgres

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/usermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/persistence/postgres/psqlmodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type gameRepo struct {
	gormDb *gorm.DB
}

func NewGameRepo() (gamemodel.Repo, error) {
	gormDb, err := NewSession()
	if err != nil {
		return nil, err
	}
	return &gameRepo{gormDb: gormDb}, nil
}

func (repo *gameRepo) Get(id gamemodel.GameIdVo) (*gamemodel.GameAgg, error) {
	gameModel := psqlmodel.WorldModel{Id: id.Uuid()}
	result := repo.gormDb.First(&gameModel)
	if result.Error != nil {
		return nil, result.Error
	}

	game := gameModel.ToAggregate()
	return &game, nil
}

func (repo *gameRepo) GetByUserId(userId usermodel.UserIdVo) (*gamemodel.GameAgg, error) {
	gameModel := psqlmodel.WorldModel{UserId: userId.Uuid()}
	result := repo.gormDb.First(&gameModel)
	if result.Error != nil {
		return nil, result.Error
	}

	game := gameModel.ToAggregate()
	return &game, nil
}

func (repo *gameRepo) Update(game gamemodel.GameAgg) error {
	return nil
}

func (repo *gameRepo) GetAll() ([]gamemodel.GameAgg, error) {
	var gameModels []psqlmodel.WorldModel
	result := repo.gormDb.Select("Id", "Width", "Height", "CreatedAt", "UpdatedAt").Find(&gameModels)
	if result.Error != nil {
		return nil, result.Error
	}

	gameAggregates := lo.Map(gameModels, func(model psqlmodel.WorldModel, _ int) gamemodel.GameAgg {
		return model.ToAggregate()
	})

	return gameAggregates, nil
}

func (repo *gameRepo) Add(game gamemodel.GameAgg) error {
	gameModel := psqlmodel.NewWorldModel(game)
	res := repo.gormDb.Create(&gameModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *gameRepo) ReadLockAccess(gameId gamemodel.GameIdVo) func() {
	return func() {}
}

func (repo *gameRepo) LockAccess(gameId gamemodel.GameIdVo) func() {
	return func() {}
}
