package service

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	commonapplicationevent "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/event"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/application/notification"
)

type LiveGameApplicationService interface {
	CreateLiveGame(gameId gamemodel.GameId) error
	ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) error
	AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
	RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
	AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error
	RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
}

type LiveGameApplicationServe struct {
	liveGameRepository    livegamemodel.LiveGameRepository
	gameRepository        gamemodel.GameRepository
	notificationPublisher commonnotification.NotificationPublisher
}

func NewLiveGameApplicationService(
	liveGameRepository livegamemodel.LiveGameRepository,
	gameRepository gamemodel.GameRepository,
	notificationPublisher commonnotification.NotificationPublisher,
) *LiveGameApplicationServe {
	return &LiveGameApplicationServe{
		liveGameRepository:    liveGameRepository,
		gameRepository:        gameRepository,
		notificationPublisher: notificationPublisher,
	}
}

func (serve *LiveGameApplicationServe) CreateLiveGame(gameId gamemodel.GameId) error {
	game, err := serve.gameRepository.Get(gameId)
	if err != nil {
		return err
	}
	newLiveGame := livegamemodel.NewLiveGame(livegamemodel.NewLiveGameId(gameId.GetId()), game.GetUnitBlock())
	serve.liveGameRepository.Add(newLiveGame)

	return nil
}

func (serve *LiveGameApplicationServe) ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) error {
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
			commonapplicationevent.NewZoomedAreaUpdatedApplicationEventChannel(liveGameId, playerId),
			commonapplicationevent.NewZoomedAreaUpdatedApplicationEvent(liveGameId, playerId, area, unitBlock),
		)
	}
	return nil
}

func (serve *LiveGameApplicationServe) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
	unlocker := serve.liveGameRepository.LockAccess(liveGameId)
	defer unlocker()

	liveGame, err := serve.liveGameRepository.Get(liveGameId)
	if err != nil {
		return err
	}

	liveGame.AddPlayer(playerId)
	serve.liveGameRepository.Update(liveGameId, liveGame)

	serve.notificationPublisher.Publish(
		commonapplicationevent.NewGameInfoUpdatedApplicationEventChannel(liveGameId, playerId),
		commonapplicationevent.NewGameInfoUpdatedApplicationEvent(liveGameId, playerId, liveGame.GetDimension()),
	)

	return nil
}

func (serve *LiveGameApplicationServe) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
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

func (serve *LiveGameApplicationServe) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error {
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
		commonapplicationevent.NewAreaZoomedApplicationEventChannel(liveGameId, playerId),
		commonapplicationevent.NewAreaZoomedApplicationEvent(liveGameId, playerId, area, unitBlock),
	)

	return nil
}

func (serve *LiveGameApplicationServe) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
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
