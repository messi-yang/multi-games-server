package pgrepo

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/pgmodel"
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
		sharedkernelmodel.NewUserIdVo(gamerModel.UserId),
	)
}

type gamerRepo struct {
	dbClient *gorm.DB
}

func NewGamerRepo() (repository gamermodel.Repo, err error) {
	dbClient, err := pgmodel.NewClient()
	if err != nil {
		return repository, err
	}
	return &gamerRepo{dbClient: dbClient}, nil
}

func (repo *gamerRepo) Add(gamer gamermodel.GamerAgg) error {
	gamerModel := newGamerModel(gamer)
	res := repo.dbClient.Create(&gamerModel)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *gamerRepo) Get(gamerId commonmodel.GamerIdVo) (gamer gamermodel.GamerAgg, err error) {
	gamerModel := pgmodel.GamerModel{Id: gamerId.Uuid()}
	result := repo.dbClient.First(&gamerModel)
	if result.Error != nil {
		return gamer, result.Error
	}
	return parseGamerModel(gamerModel), nil
}

func (repo *gamerRepo) FindGamerByUserId(userId sharedkernelmodel.UserIdVo) (gamer gamermodel.GamerAgg, gamerFound bool, err error) {
	gamerModels := []pgmodel.GamerModel{}
	result := repo.dbClient.Find(&gamerModels, pgmodel.GamerModel{UserId: userId.Uuid()})
	if result.Error != nil {
		return gamer, gamerFound, result.Error
	}
	gamerFound = result.RowsAffected >= 1
	if !gamerFound {
		return gamer, false, nil
	}
	return parseGamerModel(gamerModels[0]), true, nil
}

func (repo *gamerRepo) GetAll() (gamers []gamermodel.GamerAgg, err error) {
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
