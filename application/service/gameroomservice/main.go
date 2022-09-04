package gameroomservice

import (
	"errors"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitmapdto"
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
	GetUnitMapByArea(gameId uuid.UUID, areaDto areadto.Dto) (unitMapDto unitmapdto.Dto, err error)
	TcikAllUnitMaps() error
	ReviveUnits(gameId uuid.UUID, coordinateDtos []coordinatedto.Dto) error
	GetUnitMapSize(gameId uuid.UUID) (mapsizedto.Dto, error)
	GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDtos []coordinatedto.Dto, areaDto areadto.Dto) ([]coordinatedto.Dto, []unitdto.Dto, error)
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

func (grs *serviceImplement) GetUnitMapByArea(gameId uuid.UUID, areaDto areadto.Dto) (unitmapdto.Dto, error) {
	area, err := areadto.FromDto(areaDto)
	if err != nil {
		return unitmapdto.Dto{}, err
	}
	unitMap, err := grs.gameRoomDomainService.GetUnitMapByArea(gameId, area)
	if err != nil {
		return unitmapdto.Dto{}, err
	}

	unitMapDto := unitmapdto.ToDto(unitMap)

	return unitMapDto, nil
}

func (grs *serviceImplement) TcikAllUnitMaps() error {
	if grs.gameUnitMapTickedEvent == nil {
		return ErrEventNotFound
	}

	gameRooms := grs.gameRoomDomainService.GetAllRooms()
	for _, gameRoom := range gameRooms {
		updatedAt, err := grs.gameRoomDomainService.TickUnitMap(gameRoom.GetGameId())
		if err != nil {
			continue
		}
		grs.gameUnitMapTickedEvent.Publish(gameRoom.GetGameId(), updatedAt)
	}

	return nil
}

func (grs *serviceImplement) GetUnitMapSize(gameId uuid.UUID) (mapsizedto.Dto, error) {
	gameRoom, err := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return mapsizedto.Dto{}, err
	}
	return mapsizedto.ToDto(gameRoom.GetUnitMapSize()), nil
}

func (grs *serviceImplement) GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDtos []coordinatedto.Dto, areaDto areadto.Dto) ([]coordinatedto.Dto, []unitdto.Dto, error) {
	gameRoom, err := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return nil, nil, err
	}

	coordinates, err := coordinatedto.FromDtoList(coordinateDtos)
	if err != nil {
		return nil, nil, err
	}

	area, err := areadto.FromDto(areaDto)
	if err != nil {
		return nil, nil, err
	}

	coordinatesInArea := area.FilterCoordinates(coordinates)
	units, err := gameRoom.GetUnitsWithCoordinates(coordinatesInArea)
	if err != nil {
		return nil, nil, err
	}

	return coordinatedto.ToDtoList(coordinatesInArea), unitdto.ToDtoList(units), nil
}

func (grs *serviceImplement) ReviveUnits(gameId uuid.UUID, coordinateDtos []coordinatedto.Dto) error {
	if grs.gameUnitsRevivedEvent == nil {
		return ErrEventNotFound
	}

	coordinates, err := coordinatedto.FromDtoList(coordinateDtos)
	if err != nil {
		return err
	}

	revivedAt, err := grs.gameRoomDomainService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return err
	}

	grs.gameUnitsRevivedEvent.Publish(gameId, coordinateDtos, revivedAt)

	return nil
}
