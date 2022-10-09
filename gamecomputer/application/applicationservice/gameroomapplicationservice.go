package applicationservice

import (
	"errors"
	"fmt"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/integrationevent"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type GameRoomApplicationService interface {
	CreateGameRoom(width int, height int) (gameId uuid.UUID, err error)

	AddPlayerToGameRoom(gameId uuid.UUID, player entity.Player) error
	RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) error

	AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error
	RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) error

	TcikUnitMapInAllGames()
	ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) error
}

type gameRoomApplicationServiceImplement struct {
	gameRoomDomainService domainservice.GameRoomDomainService
	integrationEventBus   eventbus.IntegrationEventBus
}

type GameRoomApplicationServiceConfiguration struct {
	GameRoomDomainService domainservice.GameRoomDomainService
	IntegrationEventBus   eventbus.IntegrationEventBus
}

func NewGameRoomApplicationService(config GameRoomApplicationServiceConfiguration) GameRoomApplicationService {
	return &gameRoomApplicationServiceImplement{
		gameRoomDomainService: config.GameRoomDomainService,
		integrationEventBus:   config.IntegrationEventBus,
	}
}

func (grs *gameRoomApplicationServiceImplement) CreateGameRoom(width int, height int) (gameId uuid.UUID, err error) {
	mapSize, err := valueobject.NewMapSize(width, height)
	if err != nil {
		return uuid.UUID{}, err
	}

	newGameRoom, err := grs.gameRoomDomainService.CreateGameRoom(mapSize)
	if err != nil {
		return uuid.UUID{}, err
	}
	return newGameRoom.GetId(), nil
}

func (grs *gameRoomApplicationServiceImplement) TcikUnitMapInAllGames() {
	gameRooms := grs.gameRoomDomainService.GetAllGameRooms()
	for _, gameRoom := range gameRooms {
		updatedGameRoom, err := grs.gameRoomDomainService.TickUnitMapInGame(gameRoom.GetId())
		if err != nil {
			continue
		}

		for playerId, area := range updatedGameRoom.GetZoomedAreas() {
			unitMap, err := updatedGameRoom.GetUnitMapByArea(area)
			if err != nil {
				continue
			}
			grs.integrationEventBus.Publish(
				integrationevent.NewZoomedAreaUpdatedIntegrationEventTopic(updatedGameRoom.GetId(), playerId),
				integrationevent.NewZoomedAreaUpdatedIntegrationEvent(area, *unitMap),
			)
		}
	}
}

func (grs *gameRoomApplicationServiceImplement) ReviveUnitsInGame(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	updatedGameRoom, err := grs.gameRoomDomainService.ReviveUnitsInGame(gameId, coordinates)
	if err != nil {
		return err
	}

	for playerId, area := range updatedGameRoom.GetZoomedAreas() {
		coordinatesInArea := area.FilterCoordinates(coordinates)
		if len(coordinatesInArea) == 0 {
			continue
		}
		unitMap, err := updatedGameRoom.GetUnitMapByArea(area)
		if err != nil {
			continue
		}
		grs.integrationEventBus.Publish(
			integrationevent.NewZoomedAreaUpdatedIntegrationEventTopic(updatedGameRoom.GetId(), playerId),
			integrationevent.NewZoomedAreaUpdatedIntegrationEvent(area, *unitMap),
		)
	}
	return nil
}

func (grs *gameRoomApplicationServiceImplement) AddPlayerToGameRoom(gameId uuid.UUID, player entity.Player) error {
	updatedGameRoom, err := grs.gameRoomDomainService.AddPlayerToGameRoom(gameId, player)
	if err != nil {
		fmt.Println(err)
		return err
	}

	grs.integrationEventBus.Publish(
		integrationevent.NewGameInfoUpdatedIntegrationEventTopic(gameId, player.GetId()),
		integrationevent.NewGameInfoUpdatedIntegrationEvent(updatedGameRoom.GetMapSize()),
	)

	return nil
}

func (grs *gameRoomApplicationServiceImplement) RemovePlayerFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) error {
	_, err := grs.gameRoomDomainService.RemovePlayerFromGameRoom(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (grs *gameRoomApplicationServiceImplement) AddZoomedAreaToGameRoom(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	updatedGameRoom, err := grs.gameRoomDomainService.AddZoomedAreaToGameRoom(gameId, playerId, area)
	if err != nil {
		return err
	}

	unitMap, err := updatedGameRoom.GetUnitMapByArea(area)
	if err != nil {
		return err
	}

	grs.integrationEventBus.Publish(
		integrationevent.NewAreaZoomedIntegrationEventTopic(gameId, playerId),
		integrationevent.NewAreaZoomedIntegrationEvent(area, *unitMap),
	)

	return nil
}

func (grs *gameRoomApplicationServiceImplement) RemoveZoomedAreaFromGameRoom(gameId uuid.UUID, playerId uuid.UUID) error {
	_, err := grs.gameRoomDomainService.RemoveZoomedAreaFromGameRoom(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
