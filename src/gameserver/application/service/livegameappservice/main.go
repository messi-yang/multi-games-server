package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/areazoomedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/gameinfoupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/intgrevent/zoomedareaupdatedintgrevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/messaging/intgreventpublisher"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/areaviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/coordinateviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/dimensionviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/viewmodel/unitblockviewmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/playermodel"
)

type Service interface {
	CreateLiveGame(rawGameId string)
	BuildItemInLiveGame(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel, rawItemId string)
	DestroyItemInLiveGame(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel)
	AddPlayerToLiveGame(rawLiveGameId string, rawPlayerId string)
	RemovePlayerFromLiveGame(rawLiveGameId string, rawPlayerId string)
	AddZoomedAreaToLiveGame(rawLiveGameId string, rawPlayerId string, rawArea areaviewmodel.ViewModel)
	RemoveZoomedAreaFromLiveGame(rawLiveGameId string, rawPlayerId string)
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

func (serve *serve) publishZoomedAreaUpdatedEvents(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) error {
	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	for playerId, area := range liveGame.GetZoomedAreas() {
		if !area.IncludesAnyCoordinates([]commonmodel.Coordinate{coordinate}) {
			continue
		}
		unitBlock, err := liveGame.GetUnitBlockByArea(area)
		if err != nil {
			continue
		}
		serve.intgrEventPublisher.Publish(
			intgrevent.CreateLiveGameClientChannel(liveGameId.ToString(), playerId.ToString()),
			zoomedareaupdatedintgrevent.New(
				liveGameId.ToString(),
				playerId.ToString(),
				areaviewmodel.New(area),
				unitblockviewmodel.New(unitBlock),
			).Serialize(),
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
	newLiveGame := livegamemodel.NewLiveGame(liveGameId, game.GetUnitBlock())

	serve.liveGameRepo.Add(newLiveGame)
}

func (serve *serve) BuildItemInLiveGame(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel, rawItemId string) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	itemId, err := itemmodel.NewItemId(rawItemId)
	if err != nil {
		return
	}
	coordinate, err := rawCoordinate.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.BuildItem(coordinate, itemId)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	serve.publishZoomedAreaUpdatedEvents(liveGameId, coordinate)
}

func (serve *serve) DestroyItemInLiveGame(rawLiveGameId string, rawCoordinate coordinateviewmodel.ViewModel) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	coordinate, err := rawCoordinate.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	err = liveGame.DestroyItem(coordinate)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.publishZoomedAreaUpdatedEvents(liveGameId, coordinate)
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
		gameinfoupdatedintgrevent.New(rawLiveGameId, rawPlayerId, dimensionviewmodel.New(liveGame.GetDimension())).Serialize(),
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

func (serve *serve) AddZoomedAreaToLiveGame(rawLiveGameId string, rawPlayerId string, rawArea areaviewmodel.ViewModel) {
	liveGameId, err := livegamemodel.NewLiveGameId(rawLiveGameId)
	if err != nil {
		return
	}
	playerId, err := playermodel.NewPlayerId(rawPlayerId)
	if err != nil {
		return
	}
	area, err := rawArea.ToValueObject()
	if err != nil {
		return
	}

	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return
	}

	if err = liveGame.AddZoomedArea(playerId, area); err != nil {
		return
	}

	unitBlock, err := liveGame.GetUnitBlockByArea(area)
	if err != nil {
		return
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)
	serve.intgrEventPublisher.Publish(
		intgrevent.CreateLiveGameClientChannel(rawLiveGameId, rawPlayerId),
		areazoomedintgrevent.New(rawLiveGameId, rawPlayerId, rawArea, unitblockviewmodel.New(unitBlock)).Serialize(),
	)
}

func (serve *serve) RemoveZoomedAreaFromLiveGame(rawLiveGameId string, rawPlayerId string) {
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

	liveGame.RemoveZoomedArea(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)
}
