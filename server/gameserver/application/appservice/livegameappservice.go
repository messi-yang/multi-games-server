package appservice

import (
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/model/livegamemodel"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type LiveGameAppService interface {
	CreateLiveGame(gameId gamemodel.GameId) error
	BuildItemInLiveGame(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) error
	AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error
	RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error
	AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) error
	RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error
}

type LiveGameAppServe struct {
	liveGameRepository    livegamemodel.LiveGameRepository
	gameRepository        gamemodel.GameRepository
	notificationPublisher commonnotification.NotificationPublisher
}

func NewLiveGameAppService(
	liveGameRepository livegamemodel.LiveGameRepository,
	gameRepository gamemodel.GameRepository,
	notificationPublisher commonnotification.NotificationPublisher,
) *LiveGameAppServe {
	return &LiveGameAppServe{
		liveGameRepository:    liveGameRepository,
		gameRepository:        gameRepository,
		notificationPublisher: notificationPublisher,
	}
}

func (serve *LiveGameAppServe) CreateLiveGame(gameId gamemodel.GameId) error {
	game, err := serve.gameRepository.Get(gameId)
	if err != nil {
		return err
	}
	liveGameId, _ := livegamemodel.NewLiveGameId(gameId.GetId().String())
	newLiveGame := livegamemodel.NewLiveGame(liveGameId, game.GetUnitBlock())
	serve.liveGameRepository.Add(newLiveGame)

	return nil
}

func (serve *LiveGameAppServe) BuildItemInLiveGame(liveGameId livegamemodel.LiveGameId, coordinate commonmodel.Coordinate, itemId itemmodel.ItemId) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	err = liveGame.BuildItem(coordinate, itemId)
	if err != nil {
		return err
	}

	serve.liveGameRepository.Update(liveGameId, liveGame)

	for playerId, area := range liveGame.GetZoomedAreas() {
		if !area.IncludesAnyCoordinates([]commonmodel.Coordinate{coordinate}) {
			continue
		}
		unitBlock, err := liveGame.GetUnitBlockByArea(area)
		if err != nil {
			continue
		}
		serve.notificationPublisher.Publish(
			commonappevent.NewZoomedAreaUpdatedAppEventChannel(liveGameId, playerId),
			commonappevent.NewZoomedAreaUpdatedAppEvent(liveGameId, playerId, area, unitBlock),
		)
	}

	return nil
}

func (serve *LiveGameAppServe) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.AddPlayer(playerId)
	serve.liveGameRepository.Update(liveGameId, liveGame)

	serve.notificationPublisher.Publish(
		commonappevent.NewGameInfoUpdatedAppEventChannel(liveGameId, playerId),
		commonappevent.NewGameInfoUpdatedAppEvent(liveGameId, playerId, liveGame.GetDimension()),
	)

	return nil
}

func (serve *LiveGameAppServe) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.RemovePlayer(playerId)
	serve.liveGameRepository.Update(liveGameId, liveGame)

	return nil
}

func (serve *LiveGameAppServe) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId, area commonmodel.Area) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	if err = liveGame.AddZoomedArea(playerId, area); err != nil {
		return err
	}

	serve.liveGameRepository.Update(liveGameId, liveGame)

	unitBlock, err := liveGame.GetUnitBlockByArea(area)
	if err != nil {
		return err
	}

	serve.notificationPublisher.Publish(
		commonappevent.NewAreaZoomedAppEventChannel(liveGameId, playerId),
		commonappevent.NewAreaZoomedAppEvent(liveGameId, playerId, area, unitBlock),
	)

	return nil
}

func (serve *LiveGameAppServe) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId commonmodel.PlayerId) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.RemoveZoomedArea(playerId)
	serve.liveGameRepository.Update(liveGameId, liveGame)

	return nil
}
