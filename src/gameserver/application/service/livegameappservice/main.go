package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
)

type Service interface {
	CreateLiveGame(rawGameId string)
	BuildItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.LocationViewModel, rawItemId string)
	DestroyItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.LocationViewModel)
	AddPlayerToLiveGame(rawLiveGameId string, rawPlayerId string)
	RemovePlayerFromLiveGame(rawLiveGameId string, rawPlayerId string)
	AddObservedExtentToLiveGame(rawLiveGameId string, rawPlayerId string, rawExtent viewmodel.ExtentViewModel)
	RemoveObservedExtentFromLiveGame(rawLiveGameId string, rawPlayerId string)
}

type serve struct {
	liveGameRepo        livegamemodel.Repo
	gameRepo            gamemodel.GameRepo
	intgrEventPublisher intgrevent.IntgrEventPublisher
}

func New(
	liveGameRepo livegamemodel.Repo,
	gameRepo gamemodel.GameRepo,
	intgrEventPublisher intgrevent.IntgrEventPublisher,
) *serve {
	return &serve{
		liveGameRepo:        liveGameRepo,
		gameRepo:            gameRepo,
		intgrEventPublisher: intgrEventPublisher,
	}
}

func (serve *serve) publishObservedExtentUpdatedServerEvents(liveGameId livegamemodel.LiveGameId, location commonmodel.Location) error {
	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	for playerId, extent := range liveGame.GetObservedExtents() {
		if !extent.IncludesAnyLocations([]commonmodel.Location{location}) {
			continue
		}
		unitMap, err := liveGame.GetUnitMapByExtent(extent)
		if err != nil {
			continue
		}
		serve.intgrEventPublisher.Publish(
			intgrevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			intgrevent.Marshal(intgrevent.NewObservedExtentUpdatedEvent(
				liveGameId.ToString(),
				playerId.ToString(),
				viewmodel.NewExtentViewModel(extent),
				viewmodel.NewUnitMapViewModel(unitMap),
			)),
		)
	}

	return nil
}

func (serve *serve) CreateLiveGame(rawGameId string) {
	gameId, err := gamemodel.NewGameId(rawGameId)
	if err != nil {
		return
	}

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	liveGameId, _ := livegamemodel.NewLiveGameId(gameId.ToString())
	newLiveGame := livegamemodel.NewLiveGame(liveGameId, game.GetUnitMap())

	serve.liveGameRepo.Add(newLiveGame)
}

func (serve *serve) BuildItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.LocationViewModel, rawItemId string) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return
	}
	location, err := rawLocation.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.BuildItem(location, itemId)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	serve.publishObservedExtentUpdatedServerEvents(liveGameId, location)
}

func (serve *serve) DestroyItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.LocationViewModel) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	location, err := rawLocation.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.DestroyItem(location)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.publishObservedExtentUpdatedServerEvents(liveGameId, location)
}

func (serve *serve) AddPlayerToLiveGame(rawLiveGameId string, rawPlayerId string) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	liveGame.AddPlayer(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(rawLiveGameId, rawPlayerId),
		intgrevent.Marshal(
			intgrevent.NewGameInfoUpdatedEvent(rawLiveGameId, rawPlayerId, viewmodel.NewMapSizeViewModel(liveGame.GetMapSize())),
		),
	)
}

func (serve *serve) RemovePlayerFromLiveGame(rawLiveGameId string, rawPlayerId string) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	liveGame.RemovePlayer(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
}

func (serve *serve) AddObservedExtentToLiveGame(rawLiveGameId string, rawPlayerId string, rawExtent viewmodel.ExtentViewModel) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
	if err != nil {
		return
	}
	extent, err := rawExtent.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	if err = liveGame.AddObservedExtent(playerId, extent); err != nil {
		return
	}

	unitMap, err := liveGame.GetUnitMapByExtent(extent)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(rawLiveGameId, rawPlayerId),
		intgrevent.Marshal(
			intgrevent.NewExtentObservedEvent(rawLiveGameId, rawPlayerId, rawExtent, viewmodel.NewUnitMapViewModel(unitMap)),
		),
	)
}

func (serve *serve) RemoveObservedExtentFromLiveGame(rawLiveGameId string, rawPlayerId string) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	liveGame.RemoveObservedExtent(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
}
