package pgrepository

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/common/infrastructure/pgmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func newGamerModel(user gamermodel.GamerAgg) pgmodel.GamerModel {
	return pgmodel.GamerModel{
		Id:     user.GetId().Uuid(),
		UserId: user.GetUserId().Uuid(),
	}
}

func parseGamerModel(gamerModel pgmodel.GamerModel) gamermodel.GamerAgg {
	return gamermodel.NewGamerAgg(
		commonmodel.NewGamerIdVo(gamerModel.Id),
		commonmodel.NewUserIdVo(gamerModel.UserId),
	)
}

type gamerRepository struct {
	dbClient *gorm.DB
}

func NewGamerRepository() (repository gamermodel.Repository, err error) {
	dbClient, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &gamerRepository{dbClient: dbClient}, nil
}

func (repo *gamerRepository) GetAll() (gamers []gamermodel.GamerAgg, err error) {
	var gamerModels []pgmodel.GamerModel
	result := repo.dbClient.Find(&gamerModels)
	if result.Error != nil {
		err = result.Error
		return gamers, err
	}

	gamers = lo.Map(gamerModels, func(gamerModel pgmodel.GamerModel, _ int) gamermodel.GamerAgg {
		return parseGamerModel(gamerModel)
	})
	return gamers, nil
}
