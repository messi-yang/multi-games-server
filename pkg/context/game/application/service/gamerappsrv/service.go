package gamerappsrv

import (
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/domain/model/sharedkernelmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetGamerByUserId(GetGamerByUserIdQuery) (jsondto.GamerAggDto, error)
	CreateGamer(CreateGamerCommand) (gamerIdDto uuid.UUID, err error)
	QueryGamers(QueryGamersQuery) ([]jsondto.GamerAggDto, error)
}

type serve struct {
	gamerRepo gamermodel.Repo
}

func NewService(gamerRepo gamermodel.Repo) Service {
	return &serve{
		gamerRepo: gamerRepo,
	}
}

func (serve *serve) GetGamerByUserId(query GetGamerByUserIdQuery) (gamerDto jsondto.GamerAggDto, err error) {
	userId := sharedkernelmodel.NewUserIdVo(query.UserId)
	gamer, err := serve.gamerRepo.GetGamerByUserId(userId)
	if err != nil {
		return gamerDto, err
	}

	return jsondto.NewGamerAggDto(gamer), nil
}

func (serve *serve) CreateGamer(command CreateGamerCommand) (gamerIdDto uuid.UUID, err error) {
	userId := sharedkernelmodel.NewUserIdVo(command.UserId)
	_, gamerFound, err := serve.gamerRepo.FindGamerByUserId(userId)
	if err != nil {
		return gamerIdDto, err
	}
	if gamerFound {
		return gamerIdDto, fmt.Errorf("already has a gamer with userId of %s", userId.Uuid().String())
	}
	newGamer := gamermodel.NewGamerAgg(commonmodel.NewGamerIdVo(uuid.New()), userId)
	err = serve.gamerRepo.Add(newGamer)
	if err != nil {
		return gamerIdDto, err
	}
	return newGamer.GetId().Uuid(), nil
}

func (serve *serve) QueryGamers(query QueryGamersQuery) (itemDtos []jsondto.GamerAggDto, err error) {
	gamers, err := serve.gamerRepo.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(gamers, func(gamer gamermodel.GamerAgg, _ int) jsondto.GamerAggDto {
		return jsondto.NewGamerAggDto(gamer)
	}), nil
}
