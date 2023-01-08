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
	BuildItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.Location, rawItemId string)
	DestroyItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.Location)
	AddPlayerToLiveGame(rawLiveGameId string, rawPlayerId string)
	RemovePlayerFromLiveGame(rawLiveGameId string, rawPlayerId string)
	AddObservedRangeToLiveGame(rawLiveGameId string, rawPlayerId string, rangeVm viewmodel.Range)
	RemoveObservedRangeFromLiveGame(rawLiveGameId string, rawPlayerId string)
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
		unitMap, err := liveGame.GetUnitMapByRange(rangeVo)
		if err != nil {
			continue
		}
		serve.intgrEventPublisher.Publish(
			intgrevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			intgrevent.Marshal(intgrevent.NewObservedRangeUpdatedEvent(
				liveGameId.ToString(),
				playerId.ToString(),
				viewmodel.NewRange(rangeVo),
				viewmodel.NewUnitMap(unitMap),
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

func (serve *serve) BuildItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.Location, rawItemId string) {
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

	serve.publishObservedRangeUpdatedServerEvents(liveGameId, location)
}

func (serve *serve) DestroyItemInLiveGame(rawLiveGameId string, rawLocation viewmodel.Location) {
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
	serve.publishObservedRangeUpdatedServerEvents(liveGameId, location)
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
			intgrevent.NewGameInfoUpdatedEvent(rawLiveGameId, rawPlayerId, viewmodel.NewMapSize(liveGame.GetMapSize())),
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

func (serve *serve) AddObservedRangeToLiveGame(rawLiveGameId string, rawPlayerId string, rangeVm viewmodel.Range) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
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

	unitMap, err := liveGame.GetUnitMapByRange(rangeVo)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(rawLiveGameId, rawPlayerId),
		intgrevent.Marshal(
			intgrevent.NewRangeObservedEvent(rawLiveGameId, rawPlayerId, rangeVm, viewmodel.NewUnitMap(unitMap)),
		),
	)
}

func (serve *serve) RemoveObservedRangeFromLiveGame(rawLiveGameId string, rawPlayerId string) {
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

	liveGame.RemoveObservedRange(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
}
