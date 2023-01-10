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
	CreateLiveGame(gameIdVm string)
	BuildItemInLiveGame(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string)
	DestroyItemInLiveGame(liveGameIdVm string, locationVm viewmodel.LocationVm)
	AddPlayerToLiveGame(liveGameIdVm string, playerIdVm string)
	RemovePlayerFromLiveGame(liveGameIdVm string, playerIdVm string)
	AddObservedRangeToLiveGame(liveGameIdVm string, playerIdVm string, rangeVm viewmodel.RangeVm)
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

func (serve *serve) publishObservedRangeUpdatedServerEvents(liveGameId livegamemodel.LiveGameId, location commonmodel.Location) error {
	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	for playerId, rangeVo := range liveGame.GetObservedRanges() {
		if !rangeVo.IncludesAnyLocations([]commonmodel.Location{location}) {
			continue
		}
		map_, err := liveGame.GetMapByRange(rangeVo)
		if err != nil {
			continue
		}
		serve.intgrEventPublisher.Publish(
			intgrevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			intgrevent.Marshal(intgrevent.NewObservedRangeUpdatedIntgrEvent(
				liveGameId.ToString(),
				playerId.ToString(),
				viewmodel.NewRangeVm(rangeVo),
				viewmodel.NewMapVm(map_),
			)),
		)
	}

	return nil
}

func (serve *serve) CreateLiveGame(gameIdVm string) {
	gameId, err := gamemodel.NewGameId(gameIdVm)
	if err != nil {
		return
	}

	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return
	}

	liveGameId, _ := livegamemodel.NewLiveGameId(gameId.ToString())
	newLiveGame := livegamemodel.NewLiveGame(liveGameId, game.GetMap())

	serve.liveGameRepo.Add(newLiveGame)
}

func (serve *serve) BuildItemInLiveGame(liveGameIdVm string, locationVm viewmodel.LocationVm, itemIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameId(liveGameIdVm)
	if err != nil {
		return
	}
	itemId, err := itemmodel.NewItemId(itemIdVm)
	if err != nil {
		return
	}
	location, err := locationVm.ToValueObject()
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

	serve.publishObservedRangeUpdatedServerEvents(liveGameId, location)
}

func (serve *serve) DestroyItemInLiveGame(liveGameIdVm string, locationVm viewmodel.LocationVm) {
	liveGameId, err := livegamemodel.NewLiveGameId(liveGameIdVm)
	if err != nil {
		return
	}
	location, err := locationVm.ToValueObject()
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
	serve.publishObservedRangeUpdatedServerEvents(liveGameId, location)
}

func (serve *serve) AddPlayerToLiveGame(liveGameIdVm string, playerIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameId(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(playerIdVm)
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

	view, _ := liveGame.GetPlayerView(playerId)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		intgrevent.Marshal(
			intgrevent.NewGameJoinedIntgrEvent(liveGameIdVm, playerIdVm, viewmodel.NewViewVm(view), viewmodel.NewDimensionVm(liveGame.GetDimension())),
		),
	)
}

func (serve *serve) RemovePlayerFromLiveGame(liveGameIdVm string, playerIdVm string) {
	liveGameId, err := livegamemodel.NewLiveGameId(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(playerIdVm)
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

func (serve *serve) AddObservedRangeToLiveGame(liveGameIdVm string, playerIdVm string, rangeVm viewmodel.RangeVm) {
	liveGameId, err := livegamemodel.NewLiveGameId(liveGameIdVm)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(playerIdVm)
	if err != nil {
		return
	}
	rangeVo, err := rangeVm.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	if err = liveGame.AddObservedRange(playerId, rangeVo); err != nil {
		return
	}

	map_, err := liveGame.GetMapByRange(rangeVo)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(liveGameIdVm, playerIdVm),
		intgrevent.Marshal(
			intgrevent.NewRangeObservedIntgrEvent(liveGameIdVm, playerIdVm, rangeVm, viewmodel.NewMapVm(map_)),
		),
	)
}
