package applicationservice

import (
	"errors"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/sandbox"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/service/gameservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type GameApplicationService interface {
	CreateGame(game sandbox.Sandbox) (err error)

	AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) error
	RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) error

	AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error
	RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) error

	ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) error
}

type gameApplicationServiceImplement struct {
	gameDomainService   gameservice.GameDomainService
	integrationEventBus eventbus.IntegrationEventBus
}

type GameApplicationServiceConfiguration struct {
	GameDomainService   gameservice.GameDomainService
	IntegrationEventBus eventbus.IntegrationEventBus
}

func NewGameApplicationService(config GameApplicationServiceConfiguration) GameApplicationService {
	return &gameApplicationServiceImplement{
		gameDomainService:   config.GameDomainService,
		integrationEventBus: config.IntegrationEventBus,
	}
}

func (grs *gameApplicationServiceImplement) CreateGame(game sandbox.Sandbox) error {
	err := grs.gameDomainService.CreateGame(game)
	if err != nil {
		return err
	}

	return nil
}

func (grs *gameApplicationServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	updatedGame, err := grs.gameDomainService.ReviveUnitsInGame(gameId, coordinates)
	if err != nil {
		return err
	}

	for playerId, area := range updatedGame.GetZoomedAreas() {
		coordinatesInArea := area.FilterCoordinates(coordinates)
		if len(coordinatesInArea) == 0 {
			continue
		}
		unitMap, err := updatedGame.GetUnitMapByArea(area)
		if err != nil {
			continue
		}
		grs.integrationEventBus.Publish(
			integrationevent.NewZoomedAreaUpdatedIntegrationEventTopic(updatedGame.GetId(), playerId),
			integrationevent.NewZoomedAreaUpdatedIntegrationEvent(area, unitMap),
		)
	}
	return nil
}

func (grs *gameApplicationServiceImplement) AddPlayerToGame(gameId uuid.UUID, playerId uuid.UUID) error {
	updatedGame, err := grs.gameDomainService.AddPlayerToGame(gameId, playerId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	grs.integrationEventBus.Publish(
		integrationevent.NewGameInfoUpdatedIntegrationEventTopic(gameId, playerId),
		integrationevent.NewGameInfoUpdatedIntegrationEvent(updatedGame.GetMapSize()),
	)

	return nil
}

func (grs *gameApplicationServiceImplement) RemovePlayerFromGame(gameId uuid.UUID, playerId uuid.UUID) error {
	_, err := grs.gameDomainService.RemovePlayerFromGame(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (grs *gameApplicationServiceImplement) AddZoomedAreaToGame(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	updatedGame, err := grs.gameDomainService.AddZoomedAreaToGame(gameId, playerId, area)
	if err != nil {
		return err
	}

	unitMap, err := updatedGame.GetUnitMapByArea(area)
	if err != nil {
		return err
	}

	grs.integrationEventBus.Publish(
		integrationevent.NewAreaZoomedIntegrationEventTopic(gameId, playerId),
		integrationevent.NewAreaZoomedIntegrationEvent(area, unitMap),
	)

	return nil
}

func (grs *gameApplicationServiceImplement) RemoveZoomedAreaFromGame(gameId uuid.UUID, playerId uuid.UUID) error {
	_, err := grs.gameDomainService.RemoveZoomedAreaFromGame(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
