package playerappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/samber/lo"
)

type Service interface {
	GetAllPlayers(presenter Presenter)
}

type serve struct {
	playerRepo playermodel.Repo
}

func New(playerRepo playermodel.Repo) Service {
	return &serve{playerRepo: playerRepo}
}

func (serve *serve) GetAllPlayers(presenter Presenter) {
	players := serve.playerRepo.GetAll()
	itemCameraModels := lo.Map(players, func(player playermodel.PlayerAgg, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(player)
	})
	presenter.OnSuccess(itemCameraModels)
}
