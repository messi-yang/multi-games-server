package gameroomservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmaptickedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitsrevivedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type Service interface {
	CreateRoom(width int, height int) (gameId uuid.UUID, err error)
	GetUnitMapByArea(gameId uuid.UUID, area valueobject.Area) (unitMap valueobject.UnitMap, err error)
	TcikAllUnitMaps() error
	ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error
	GetUnitMapSize(gameId uuid.UUID) (valueobject.MapSize, error)
	AreaIncludesAnyCoordinates(area valueobject.Area, coordinates []valueobject.Coordinate) (bool, error)
}

type serviceImplement struct {
	gameRoomDomainService  gameroomservice.Service
	gameUnitMapTickedEvent gameunitmaptickedevent.Event
	gameUnitsRevivedEvent  gameunitsrevivedevent.Event
}

type Configuration struct {
	GameRoomRepository     gameroomrepository.Repository
	GameUnitMapTickedEvent gameunitmaptickedevent.Event
	UnitsRevivedEvent      gameunitsrevivedevent.Event
}

func NewService(config Configuration) Service {
	return &serviceImplement{
		gameRoomDomainService:  gameroomservice.NewService(config.GameRoomRepository),
		gameUnitMapTickedEvent: config.GameUnitMapTickedEvent,
		gameUnitsRevivedEvent:  config.UnitsRevivedEvent,
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

func (grs *serviceImplement) TcikAllUnitMaps() error {
	if grs.gameUnitMapTickedEvent == nil {
		return ErrEventNotFound
	}

	gameRooms := grs.gameRoomDomainService.GetAllRooms()
	for _, gameRoom := range gameRooms {
		err := grs.gameRoomDomainService.TickUnitMap(gameRoom.GetGameId())
		if err != nil {
			continue
		}
		grs.gameUnitMapTickedEvent.Publish(gameRoom.GetGameId())
	}

	return nil
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

func (grs *serviceImplement) ReviveUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate) error {
	if grs.gameUnitsRevivedEvent == nil {
		return ErrEventNotFound
	}

	err := grs.gameRoomDomainService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return err
	}

	grs.gameUnitsRevivedEvent.Publish(gameId, coordinates)

	return nil
}
