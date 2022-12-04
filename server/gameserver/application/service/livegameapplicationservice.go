package service

import (
	gamecommonmodel "github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/common"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/gamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/model/livegamemodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/gamedomain/service/livegameservice"
	commonnotification "github.com/dum-dum-genius/game-of-liberty-computer/server/common/notification"
	commonredis "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis"
	commonredisdto "github.com/dum-dum-genius/game-of-liberty-computer/server/common/port/adapter/notification/redis/dto"
)

type LiveGameApplicationService interface {
	CreateLiveGame(gameId gamemodel.GameId) (livegamemodel.LiveGameId, error)
	ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) error
	AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
	RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
	AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error
	RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error
}

type LiveGameApplicationServe struct {
	liveGameService       livegameservice.LiveGameService
	gameService           gameservice.GameService
	notificationPublisher commonnotification.NotificationPublisher
}

type liveGameApplicationServiceConfiguration func(serve *LiveGameApplicationServe) error

func NewLiveGameApplicationService(cfgs ...liveGameApplicationServiceConfiguration) (*LiveGameApplicationServe, error) {
	serve := &LiveGameApplicationServe{}
	for _, cfg := range cfgs {
		err := cfg(serve)
		if err != nil {
			return nil, err
		}
	}
	return serve, nil
}

func WithLiveGameService() liveGameApplicationServiceConfiguration {
	return func(serve *LiveGameApplicationServe) error {
		liveGameService, _ := livegameservice.NewLiveGameService(
			livegameservice.WithGameMemoryRepository(),
		)
		serve.liveGameService = liveGameService
		return nil
	}
}

func WithGameService() liveGameApplicationServiceConfiguration {
	return func(serve *LiveGameApplicationServe) error {
		gameService, _ := gameservice.NewGameService(
			gameservice.WithPostgresGameRepository(),
		)
		serve.gameService = gameService
		return nil
	}
}

func WithRedisNotificationPublisher() liveGameApplicationServiceConfiguration {
	return func(serve *LiveGameApplicationServe) error {
		serve.notificationPublisher = commonredis.NewRedisNotificationPublisher()
		return nil
	}
}

func (serve *LiveGameApplicationServe) CreateLiveGame(gameId gamemodel.GameId) (livegamemodel.LiveGameId, error) {
	game, err := serve.gameService.GetGame(gameId)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}
	liveGameId, err := serve.liveGameService.CreateLiveGame(game)
	if err != nil {
		return livegamemodel.LiveGameId{}, err
	}

	return liveGameId, nil
}

func (serve *LiveGameApplicationServe) ReviveUnitsInLiveGame(liveGameId livegamemodel.LiveGameId, coordinates []gamecommonmodel.Coordinate) error {
	updatedGame, err := serve.liveGameService.ReviveUnitsInLiveGame(liveGameId, coordinates)
	if err != nil {
		return err
	}

	for playerId, area := range updatedGame.GetZoomedAreas() {
		coordinatesInArea := area.FilterCoordinates(coordinates)
		if len(coordinatesInArea) == 0 {
			continue
		}
		unitBlock, err := updatedGame.GetUnitBlockByArea(area)
		if err != nil {
			continue
		}
		serve.notificationPublisher.Publish(
			commonredisdto.NewRedisZoomedAreaUpdatedEventChannel(updatedGame.GetId(), playerId),
			commonredisdto.NewRedisZoomedAreaUpdatedEvent(updatedGame.GetId(), playerId, area, unitBlock),
		)
	}
	return nil
}

func (serve *LiveGameApplicationServe) AddPlayerToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
	updatedGame, err := serve.liveGameService.AddPlayerToLiveGame(liveGameId, playerId)
	if err != nil {
		return err
	}

	serve.notificationPublisher.Publish(
		commonredisdto.NewRedisGameInfoUpdatedEventChannel(liveGameId, playerId),
		commonredisdto.NewRedisGameInfoUpdatedEvent(liveGameId, playerId, updatedGame.GetDimension()),
	)

	return nil
}

func (serve *LiveGameApplicationServe) RemovePlayerFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
	_, err := serve.liveGameService.RemovePlayerFromLiveGame(liveGameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (serve *LiveGameApplicationServe) AddZoomedAreaToLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId, area gamecommonmodel.Area) error {
	updatedGame, err := serve.liveGameService.AddZoomedAreaToLiveGame(liveGameId, playerId, area)
	if err != nil {
		return err
	}

	unitBlock, err := updatedGame.GetUnitBlockByArea(area)
	if err != nil {
		return err
	}

	serve.notificationPublisher.Publish(
		commonredisdto.NewRedisAreaZoomedEventChannel(liveGameId, playerId),
		commonredisdto.NewRedisAreaZoomedEvent(liveGameId, playerId, area, unitBlock),
	)

	return nil
}

func (serve *LiveGameApplicationServe) RemoveZoomedAreaFromLiveGame(liveGameId livegamemodel.LiveGameId, playerId gamecommonmodel.PlayerId) error {
	_, err := serve.liveGameService.RemoveZoomedAreaFromLiveGame(liveGameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
