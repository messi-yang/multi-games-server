package gameroomservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/application/eventbus"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/valueobject"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/unitmaptickedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/presenter/event/unitsrevivedevent"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type Service interface {
	CreateRoom(width int, height int) (gameId uuid.UUID, err error)

	AddPlayer(gameId uuid.UUID, player entity.Player) error
	RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error
	GetPlayers(gameId uuid.UUID) ([]entity.Player, error)

	AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error
	RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error

	GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (unitMap valueobject.UnitMap, err error)

	GetUnitMapSize(gameId uuid.UUID) (valueobject.MapSize, error)

	TcikAllUnitMaps()
	ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate)

	AreaIncludesAnyCoordinates(area valueobject.Area, coordinates []valueobject.Coordinate) (bool, error)
}

type serviceImplement struct {
	gameRoomDomainService gameroomservice.Service
	eventBus              eventbus.EventBus
}

type Configuration struct {
	GameRoomRepository gameroomrepository.Repository
	EventBus           eventbus.EventBus
}

func NewService(config Configuration) Service {
	return &serviceImplement{
		gameRoomDomainService: gameroomservice.NewService(config.GameRoomRepository),
		eventBus:              config.EventBus,
	}
}

func (grs *serviceImplement) CreateRoom(width int, height int) (gameId uuid.UUID, err error) {
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

func (grs *serviceImplement) GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (valueobject.UnitMap, error) {
	unitMap, err := grs.gameRoomDomainService.GetUnitMapByArea(gameId, area)
	if err != nil {
		return valueobject.UnitMap{}, err
	}

	return *unitMap, nil
}

func (grs *serviceImplement) TcikAllUnitMaps() {
	gameRooms := grs.gameRoomDomainService.GetAllRooms()
	for _, gameRoom := range gameRooms {
		err := grs.gameRoomDomainService.TickUnitMap(gameRoom.GetGameId())
		if err != nil {
			continue
		}
		grs.eventBus.Publish(
			unitmaptickedevent.NewEventTopic(gameRoom.GetGameId()),
			unitmaptickedevent.NewEvent(),
		)
	}
}

func (grs *serviceImplement) GetUnitMapSize(gameId uuid.UUID) (valueobject.MapSize, error) {
	gameRoom, err := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return valueobject.MapSize{}, err
	}
	return gameRoom.GetUnitMapSize(), nil
}

func (grs *serviceImplement) AreaIncludesAnyCoordinates(area valueobject.Area, coordinates []valueobject.Coordinate) (bool, error) {
	coordinatesInArea := area.FilterCoordinates(coordinates)

	return len(coordinatesInArea) > 0, nil
}

func (grs *serviceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) {
	err := grs.gameRoomDomainService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return
	}

	grs.eventBus.Publish(unitsrevivedevent.NewEventTopic(gameId), unitsrevivedevent.NewEvent(gameId, coordinates))
}

func (grs *serviceImplement) AddPlayer(gameId uuid.UUID, player entity.Player) error {
	err := grs.gameRoomDomainService.AddPlayer(gameId, player)
	if err != nil {
		return err
	}

	return nil
}

func (grs *serviceImplement) RemovePlayer(gameId uuid.UUID, playerId uuid.UUID) error {
	err := grs.gameRoomDomainService.RemovePlayer(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}

func (grs *serviceImplement) GetPlayers(gameId uuid.UUID) ([]entity.Player, error) {
	players, err := grs.gameRoomDomainService.GetAllPlayers(gameId)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (grs *serviceImplement) AddZoomedArea(gameId uuid.UUID, playerId uuid.UUID, area valueobject.Area) error {
	err := grs.gameRoomDomainService.AddZoomedArea(gameId, playerId, area)
	if err != nil {
		return err
	}

	return nil
}

func (grs *serviceImplement) RemoveZoomedArea(gameId uuid.UUID, playerId uuid.UUID) error {
	err := grs.gameRoomDomainService.RemoveZoomedArea(gameId, playerId)
	if err != nil {
		return err
	}

	return nil
}
