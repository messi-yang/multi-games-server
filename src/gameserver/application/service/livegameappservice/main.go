package livegameappservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/areazoomedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/gameinfoupdatedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/common/application/integrationevent/zoomedareaupdatedintegrationevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/src/domain/model/livegamemodel"
)

type Service interface {
	CreateLiveGame(gameId gamemodel.GameId) error
	BuildItemInLiveGame(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) error
	DestroyItemInLiveGame(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) error
	AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error
	RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error
	AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) error
	RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error
}

type serve struct {
	liveGameRepo          livegamemodel.Repo
	gameRepo              gamemodel.GameRepo
	notificationPublisher integrationevent.Publisher
}

func New(
	liveGameRepo livegamemodel.Repo,
	gameRepo gamemodel.GameRepo,
	notificationPublisher integrationevent.Publisher,
) *serve {
	return &serve{
		liveGameRepo:          liveGameRepo,
		gameRepo:              gameRepo,
		notificationPublisher: notificationPublisher,
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
		serve.notificationPublisher.Publish(
			integrationevent.CreateLiveGameClientChannel(liveGameId, playerId),
			zoomedareaupdatedintegrationevent.New(liveGameId, playerId, area, unitBlock).Serialize(),
		)
	}

	return nil
}

func (serve *serve) CreateLiveGame(gameId gamemodel.GameId) error {
	game, err := serve.gameRepo.Get(gameId)
	if err != nil {
		return err
	}
	liveGameId, _ := livegamemodel.NewLiveGameId(gameId.GetId().String())
	newLiveGame := livegamemodel.NewLiveGame(liveGameId, game.GetUnitBlock())
	serve.liveGameRepo.Add(newLiveGame)

	return nil
}

func (serve *serve) BuildItemInLiveGame(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) error {
	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	err = liveGame.BuildItem(coordinate, itemId)
	if err != nil {
		return err
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	serve.publishZoomedAreaUpdatedEvents(liveGameId, coordinate)

	return nil
}

func (serve *serve) DestroyItemInLiveGame(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate) error {
	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	err = liveGame.DestroyItem(coordinate)
	if err != nil {
		return err
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	serve.publishZoomedAreaUpdatedEvents(liveGameId, coordinate)

	return nil
}

func (serve *serve) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error {
	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.AddPlayer(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)

	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameClientChannel(liveGameId, playerId),
		gameinfoupdatedintegrationevent.New(liveGameId, playerId, liveGame.GetDimension()).Serialize(),
	)

	return nil
}

func (serve *serve) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error {
	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.RemovePlayer(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)

	return nil
}

func (serve *serve) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) error {
	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	if err = liveGame.AddZoomedArea(playerId, area); err != nil {
		return err
	}

	serve.liveGameRepo.Update(liveGameId, liveGame)

	unitBlock, err := liveGame.GetUnitBlockByArea(area)
	if err != nil {
		return err
	}

	serve.notificationPublisher.Publish(
		integrationevent.CreateLiveGameClientChannel(liveGameId, playerId),
		areazoomedintegrationevent.New(liveGameId, playerId, area, unitBlock).Serialize(),
	)

	return nil
}

func (serve *serve) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error {
	unlocker := serve.liveGameRepo.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepo.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.RemoveZoomedArea(playerId)
	serve.liveGameRepo.Update(liveGameId, liveGame)

	return nil
}
