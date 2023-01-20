package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
	"github.com/samber/lo"
)

type PlayerAppService interface {
	GetAllPlayers(presenter Presenter)
}

type playerAppServe struct {
	playerRepo playermodel.Repo
}

func NewPlayerAppService(playerRepo playermodel.Repo) PlayerAppService {
	return &playerAppServe{playerRepo: playerRepo}
}

func (playerAppServe *playerAppServe) GetAllPlayers(presenter Presenter) {
	players := playerAppServe.playerRepo.GetAll()
	itemCameraModels := lo.Map(players, func(player playermodel.PlayerAgg, _ int) viewmodel.PlayerVm {
		return viewmodel.NewPlayerVm(player)
	})
	presenter.OnSuccess(itemCameraModels)
}
