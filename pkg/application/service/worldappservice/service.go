package worldappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/dto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/samber/lo"
)

type Service interface {
	QueryWorlds()
}

type serve struct {
	worldRepository worldmodel.Repository
	presenter       Presenter
}

func NewService(worldRepository worldmodel.Repository, presenter Presenter) Service {
	return &serve{
		worldRepository: worldRepository,
		presenter:       presenter,
	}
}

func (serve *serve) QueryWorlds() {
	worlds, err := serve.worldRepository.GetAll()
	if err != nil {
		serve.presenter.OnError(err)
		return
	}
	worldDtos := lo.Map(worlds, func(world worldmodel.WorldAgg, _ int) dto.WorldAggDto {
		return dto.NewWorldAggDto(world)
	})
	responseDto := QueryWorldsResponseDto(worldDtos)
	serve.presenter.OnSuccess(responseDto)
}
