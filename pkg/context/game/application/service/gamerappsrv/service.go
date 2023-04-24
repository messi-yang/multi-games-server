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
	CreateGamer(CreateGamerCommand) (gamerIdDto uuid.UUID, err error)
	QueryGamers(QueryGamersQuery) ([]jsondto.GamerAggDto, error)
}

type serve struct {
	gamerRepository gamermodel.Repository
}

func NewService(gamerRepository gamermodel.Repository) Service {
	return &serve{
		gamerRepository: gamerRepository,
	}
}

func (serve *serve) CreateGamer(command CreateGamerCommand) (gamerIdDto uuid.UUID, err error) {
	userId := sharedkernelmodel.NewUserIdVo(command.UserId)
	_, gamerFound, err := serve.gamerRepository.FindGamerByUserId(userId)
	if err != nil {
		return gamerIdDto, err
	}
	if gamerFound {
		return gamerIdDto, fmt.Errorf("already has a gamer with userId of %s", userId.Uuid().String())
	}
	newGamer := gamermodel.NewGamerAgg(commonmodel.NewGamerIdVo(uuid.New()), userId)
	err = serve.gamerRepository.Add(newGamer)
	if err != nil {
		return gamerIdDto, err
	}
	return newGamer.GetId().Uuid(), nil
}

func (serve *serve) QueryGamers(query QueryGamersQuery) (itemDtos []jsondto.GamerAggDto, err error) {
	gamers, err := serve.gamerRepository.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(gamers, func(gamer gamermodel.GamerAgg, _ int) jsondto.GamerAggDto {
		return jsondto.NewGamerAggDto(gamer)
	}), nil
}
