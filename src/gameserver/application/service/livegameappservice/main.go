package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/messaging/intgreventpublisher"
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
	AddObservedMapRangeToLiveGame(rawLiveGameId string, rawPlayerId string, rawMapRange viewmodel.MapRangeViewModel)
	RemoveObservedMapRangeFromLiveGame(rawLiveGameId string, rawPlayerId string)
}

type serve struct {
	liveGameRepo        livegamemodel.Repo
	gameRepo            gamemodel.GameRepo
	intgrEventPublisher intgreventpublisher.Publisher
}

func New(
	liveGameRepo livegamemodel.Repo,
	gameRepo gamemodel.GameRepo,
	intgrEventPublisher intgreventpublisher.Publisher,
) *serve {
	return &serve{
		liveGameRepo:        liveGameRepo,
		gameRepo:            gameRepo,
		intgrEventPublisher: intgrEventPublisher,
	}
}

func (serve *serve) publishObservedMapRangeUpdatedServerEvents(liveGameId livegamemodel.LiveGameId, location commonmodel.Location) error {
	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	for playerId, mapRange := range liveGame.GetObservedMapRanges() {
		if !mapRange.IncludesAnyLocations([]commonmodel.Location{location}) {
			continue
		}
		unitMap, err := liveGame.GetUnitMapByMapRange(mapRange)
		if err != nil {
			continue
		}
		serve.intgrEventPublisher.Publish(
			intgrevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			intgrevent.Marshal(intgrevent.NewObservedMapRangeUpdatedEvent(
				liveGameId.ToString(),
				playerId.ToString(),
				viewmodel.NewMapRangeViewModel(mapRange),
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

	serve.publishObservedMapRangeUpdatedServerEvents(liveGameId, location)
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
	serve.publishObservedMapRangeUpdatedServerEvents(liveGameId, location)
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

func (serve *serve) AddObservedMapRangeToLiveGame(rawLiveGameId string, rawPlayerId string, rawMapRange viewmodel.MapRangeViewModel) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
	if err != nil {
		return
	}
	mapRange, err := rawMapRange.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	if err = liveGame.AddObservedMapRange(playerId, mapRange); err != nil {
		return
	}

	unitMap, err := liveGame.GetUnitMapByMapRange(mapRange)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(rawLiveGameId, rawPlayerId),
		intgrevent.Marshal(
			intgrevent.NewMapRangeObservedEvent(rawLiveGameId, rawPlayerId, rawMapRange, viewmodel.NewUnitMapViewModel(unitMap)),
		),
	)
}

func (serve *serve) RemoveObservedMapRangeFromLiveGame(rawLiveGameId string, rawPlayerId string) {
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

	liveGame.RemoveObservedMapRange(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
}
