package gamerappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/jsondto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/domain/model/gamermodel"
	"github.com/samber/lo"
)

type Service interface {
	GetGamers(GetGamersQuery) ([]jsondto.GamerAggDto, error)
}

type serve struct {
	gamerRepository gamermodel.Repository
}

func NewService(gamerRepository gamermodel.Repository) Service {
	return &serve{
		gamerRepository: gamerRepository,
	}
}

func (serve *serve) GetGamers(query GetGamersQuery) (itemDtos []jsondto.GamerAggDto, err error) {
	gamers, err := serve.gamerRepository.GetAll()
	if err != nil {
		return itemDtos, err
	}

	return lo.Map(gamers, func(gamer gamermodel.GamerAgg, _ int) jsondto.GamerAggDto {
		return jsondto.NewGamerAggDto(gamer)
	}), nil
}
