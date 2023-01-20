package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/samber/lo"
)

type GameAppService interface {
	GetAll(presenter Presenter)
}

type gameAppServe struct {
	gameRepo gamemodel.GameRepo
}

func NewGameAppService(gameRepo gamemodel.GameRepo) GameAppService {
	return &gameAppServe{gameRepo: gameRepo}
}

func (gameAppServe *gameAppServe) GetAll(presenter Presenter) {
	games, _ := gameAppServe.gameRepo.GetAll()
	gameVms := lo.Map(games, func(game gamemodel.GameAgg, _ int) viewmodel.GameVm {
		return viewmodel.NewGameVm(game)
	})
	presenter.OnSuccess(gameVms)
}
