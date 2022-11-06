package applicationservice

import (
	"errors"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/infrastructure/eventbusredis"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type GameApplicationService struct {
	gameService         *gameservice.GameService
	integrationEventBus eventbus.IntegrationEventBus
}

type gameApplicationServiceConfiguration func(service *GameApplicationService) error

func NewGameApplicationService(cfgs ...gameApplicationServiceConfiguration) (*GameApplicationService, error) {
	service := &GameApplicationService{}
	for _, cfg := range cfgs {
		err := cfg(service)
		if err != nil {
			return nil, err
		}
	}
	return service, nil
}

func WithGameService() gameApplicationServiceConfiguration {
	gameService, _ := gameservice.NewGameService(
		gameservice.WithGameMemory(),
	)
	return func(service *GameApplicationService) error {
		service.gameService = gameService
		return nil
	}
}

func WithRedisIntegrationEventBus() gameApplicationServiceConfiguration {
	return func(service *GameApplicationService) error {
		redisIntegrationEventBus, _ := eventbusredis.NewRedisIntegrationEventBus(eventbusredis.WithRedisService())
		service.integrationEventBus = redisIntegrationEventBus
		return nil
	}
}

func (grs *GameApplicationService) CreateGame(dimension valueobject.Dimension) (uuid.UUID, error) {
	gameId, err := grs.gameService.CreateGame(dimension)
	if err != nil {
		return gameId, err
	}

	return gameId, nil
}

func (grs *GameApplicationService) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	updatedGame, err := grs.gameService.ReviveUnitsInGame(gameId, coordinates)
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
		grs.integrationEventBus.Publish(
			integrationevent.NewZoomedAreaUpdatedIntegrationEventTopic(updatedGame.GetId(), playerId),
			integrationevent.NewZoomedAreaUpdatedIntegrationEvent(area, unitBlock),
		)
	}
	return nil
}

func (grs *GameApplicationService) AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) error {
	updatedGame, err := grs.gameService.AddPlayerToGame(gameId, playerId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	grs.integrationEventBus.Publish(
		integrationevent.NewGameInfoUpdatedIntegrationEventTopic(gameId, playerId),
		integrationevent.NewGameInfoUpdatedIntegrationEvent(updatedGame.GetDimension()),
	)

	return nil
}

func (grs *GameApplicationService) RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) error {
	_, err := grs.gameService.RemovePlayerFromGame(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (grs *GameApplicationService) AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	updatedGame, err := grs.gameService.AddZoomedAreaToGame(gameId, playerId, area)
	if err != nil {
		return err
	}

	unitBlock, err := updatedGame.GetUnitBlockByArea(area)
	if err != nil {
		return err
	}

	grs.integrationEventBus.Publish(
		integrationevent.NewAreaZoomedIntegrationEventTopic(gameId, playerId),
		integrationevent.NewAreaZoomedIntegrationEvent(area, unitBlock),
	)

	return nil
}

func (grs *GameApplicationService) RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) error {
	_, err := grs.gameService.RemoveZoomedAreaFromGame(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
