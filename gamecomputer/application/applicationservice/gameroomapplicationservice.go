package applicationservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/domainservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/areazoomedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/gameinfoupdatedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/zoomedareaupdatedevent"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type GameRoomApplicationService interface {
	CreateRoom(width int, height int) (gameId uuid.UUID, err error)

	AddPlayer(gameId uuid.UUID, player entity.Player) error
	RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error

	AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error
	RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error

	TcikAllUnitMaps()
	ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error
}

type gameRoomApplicationServiceImplement struct {
	gameRoomDomainService domainservice.GameRoomDomainService
	integrationEventBus   eventbus.IntegrationEventBus
}

type GameRoomApplicationServiceConfiguration struct {
	GameRoomRepository  repository.GameRoomRepository
	IntegrationEventBus eventbus.IntegrationEventBus
}

func NewGameRoomApplicationService(config GameRoomApplicationServiceConfiguration) GameRoomApplicationService {
	return &gameRoomApplicationServiceImplement{
		gameRoomDomainService: domainservice.NewGameRoomDomainService(config.GameRoomRepository),
		integrationEventBus:   config.IntegrationEventBus,
	}
}

func (grs *gameRoomApplicationServiceImplement) CreateRoom(width int, height int) (gameId uuid.UUID, err error) {
	mapSize, err := valueobject.NewMapSize(width, height)
	if err != nil {
		return uuid.UUID{}, err
	}

	gameRoom, err := grs.gameRoomDomainService.CreateGameRoom(mapSize)
	if err != nil {
		return uuid.UUID{}, err
	}
	return gameRoom.GetGameId(), nil
}

func (grs *gameRoomApplicationServiceImplement) TcikAllUnitMaps() {
	gameRooms := grs.gameRoomDomainService.GetAllRooms()
	for _, gameRoom := range gameRooms {
		err := grs.gameRoomDomainService.TickUnitMap(gameRoom.GetGameId())
		if err != nil {
			continue
		}

		for playerId, area := range gameRoom.GetZoomedAreas() {
			unitMap, err := gameRoom.GetUnitMapByArea(area)
			if err != nil {
				continue
			}
			grs.integrationEventBus.Publish(
				zoomedareaupdatedevent.NewEventTopic(gameRoom.GetGameId(), playerId),
				zoomedareaupdatedevent.NewEvent(area, *unitMap),
			)
		}
	}
}

func (grs *gameRoomApplicationServiceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	err := grs.gameRoomDomainService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return err
	}

	gameRoom, _ := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return err
	}

	for playerId, area := range gameRoom.GetZoomedAreas() {
		coordinatesInArea := area.FilterCoordinates(coordinates)
		if len(coordinatesInArea) == 0 {
			continue
		}
		unitMap, err := gameRoom.GetUnitMapByArea(area)
		if err != nil {
			continue
		}
		grs.integrationEventBus.Publish(
			zoomedareaupdatedevent.NewEventTopic(gameRoom.GetGameId(), playerId),
			zoomedareaupdatedevent.NewEvent(area, *unitMap),
		)
	}
	return nil
}

func (grs *gameRoomApplicationServiceImplement) AddPlayer(gameId uuid.UUID, player entity.Player) error {
	err := grs.gameRoomDomainService.AddPlayer(gameId, player)
	if err != nil {
		return err
	}

	gameRoom, _ := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return err
	}

	grs.integrationEventBus.Publish(
		gameinfoupdatedevent.NewEventTopic(gameId, player.GetId()),
		gameinfoupdatedevent.NewEvent(gameRoom.GetUnitMapSize()),
	)

	return nil
}

func (grs *gameRoomApplicationServiceImplement) RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error {
	err := grs.gameRoomDomainService.RemovePlayer(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (grs *gameRoomApplicationServiceImplement) AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	err := grs.gameRoomDomainService.AddZoomedArea(gameId, playerId, area)
	if err != nil {
		return err
	}

	gameRoom, _ := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return err
	}

	unitMap, err := gameRoom.GetUnitMapByArea(area)
	if err != nil {
		return err
	}

	grs.integrationEventBus.Publish(
		areazoomedevent.NewEventTopic(gameId, playerId),
		areazoomedevent.NewEvent(area, *unitMap),
	)

	return nil
}

func (grs *gameRoomApplicationServiceImplement) RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error {
	err := grs.gameRoomDomainService.RemoveZoomedArea(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
