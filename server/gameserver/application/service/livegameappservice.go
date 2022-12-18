package service

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonappevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type LiveGameAppService interface {
	CreateLiveGame(gameId gamemodel.GameId) error
	ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) error
	AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
	RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
	AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error
	RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
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

func (serve *LiveGameAppServe) ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	err = liveGame.ReviveUnits(coordinates)
	if err != nil {
		return err
	}

	serve.liveGameRepository.Update(liveGameId, liveGame)

	for playerId, area := range liveGame.GetZoomedAreas() {
		if !area.IncludesAnyCoordinates(coordinates) {
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

func (serve *LiveGameAppServe) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
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

func (serve *LiveGameAppServe) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
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

func (serve *LiveGameAppServe) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error {
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

func (serve *LiveGameAppServe) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
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
