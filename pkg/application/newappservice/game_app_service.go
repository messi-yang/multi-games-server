package newappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/intevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/service"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/library/jsonmarshaller"
)

type GameAppService interface {
	PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16)
	DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm)
}

type gameAppServe struct {
	gameRepo          gamemodel.Repo
	unitRepo          unitmodel.Repo
	itemRepo          itemmodel.Repo
	gameService       service.GameService
	IntEventPublisher intevent.IntEventPublisher
}

func NewGameAppService(
	gameRepo gamemodel.Repo,
	unitRepo unitmodel.Repo,
	itemRepo itemmodel.Repo,
	IntEventPublisher intevent.IntEventPublisher,
) GameAppService {
	return &gameAppServe{
		gameRepo:          gameRepo,
		itemRepo:          itemRepo,
		unitRepo:          unitRepo,
		gameService:       service.NewGameService(gameRepo, unitRepo, itemRepo),
		IntEventPublisher: IntEventPublisher,
	}
}

func (gameAppServe *gameAppServe) publishViewUpdatedEvents(gameId gamemodel.GameIdVo, location commonmodel.LocationVo) error {
	game, err := gameAppServe.gameRepo.Get(gameId)
	if err != nil {
		return err
	}

	for _, playerId := range game.GetPlayerIds() {
		if !game.CanPlayerSeeAnyLocations(playerId, []commonmodel.LocationVo{location}) {
			continue
		}

		// Delete this section later
		bound, _ := game.GetPlayerViewBound(playerId)
		units := gameAppServe.unitRepo.GetUnits(gameId, bound)
		view := unitmodel.NewViewVo(bound, units)
		// Delete this section later

		gameAppServe.IntEventPublisher.Publish(
			intevent.CreateGameClientChannel(gameId.ToString(), playerId.ToString()),
			jsonmarshaller.Marshal(intevent.NewViewUpdatedIntEvent(
				gameId.ToString(),
				playerId.ToString(),
				viewmodel.NewViewVm(view),
			)))
	}

	return nil
}

func (gameAppServe *gameAppServe) PlaceItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm, itemIdVm int16) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	itemId := itemmodel.NewItemIdVo(itemIdVm)
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	err = gameAppServe.gameService.PlaceItem(gameId, playerId, itemId, location)
	if err != nil {
		return
	}

	gameAppServe.publishViewUpdatedEvents(gameId, location)
}

func (gameAppServe *gameAppServe) DestroyItem(gameIdVm string, playerIdVm string, locationVm viewmodel.LocationVm) {
	gameId, err := gamemodel.NewGameIdVo(gameIdVm)
	if err != nil {
		return
	}
	playerId, err := gamemodel.NewPlayerIdVo(playerIdVm)
	if err != nil {
		return
	}
	location := commonmodel.NewLocationVo(locationVm.X, locationVm.Y)

	unlocker := gameAppServe.gameRepo.LockAccess(gameId)
	defer unlocker()

	gameAppServe.gameService.DestroyItem(gameId, playerId, location)

	gameAppServe.publishViewUpdatedEvents(gameId, location)
}
