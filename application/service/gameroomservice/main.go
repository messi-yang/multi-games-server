package gameroomservice

import (
	"errors"
	"time"

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
	GetUnitMapByArea(gameId uuid.UUID, areaDTO areadto.DTO) (unitMapDTO unitmapdto.DTO, receivedAt time.Time, err error)
	TcikAllUnitMaps() error
	ReviveUnits(gameId uuid.UUID, coordinateDTOs []coordinatedto.DTO) error
	GetUnitMapSize(gameId uuid.UUID) (mapsizedto.DTO, error)
	GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDTOs []coordinatedto.DTO, areaDTO areadto.DTO) ([]coordinatedto.DTO, []unitdto.DTO, error)
}

type serviceImplement struct {
	gameRoomDomainService  gameroomservice.GameRoomService
	gameUnitMapTickedEvent gameunitmaptickedevent.Event
	gameUnitsRevivedEvent  gameunitsrevivedevent.Event
}

type Configuration struct {
	GameRoomRepository     gameroomrepository.GameRoomRepository
	GameUnitMapTickedEvent gameunitmaptickedevent.Event
	UnitsRevivedEvent      gameunitsrevivedevent.Event
}

func NewService(config Configuration) Service {
	return &serviceImplement{
		gameRoomDomainService:  gameroomservice.NewGameRoomService(config.GameRoomRepository),
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

func (grs *serviceImplement) GetUnitMapByArea(gameId uuid.UUID, areaDTO areadto.DTO) (unitmapdto.DTO, time.Time, error) {
	area, err := areadto.FromDTO(areaDTO)
	if err != nil {
		return unitmapdto.DTO{}, time.Time{}, err
	}
	unitMap, receivedAt, err := grs.gameRoomDomainService.GetUnitMapByArea(gameId, area)
	if err != nil {
		return unitmapdto.DTO{}, time.Time{}, err
	}

	unitMapDTO := unitmapdto.ToDTO(unitMap)

	return unitMapDTO, receivedAt, nil
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

func (grs *serviceImplement) GetUnitMapSize(gameId uuid.UUID) (mapsizedto.DTO, error) {
	gameRoom, err := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return mapsizedto.DTO{}, err
	}
	return mapsizedto.ToDTO(gameRoom.GetUnitMapSize()), nil
}

func (grs *serviceImplement) GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDTOs []coordinatedto.DTO, areaDTO areadto.DTO) ([]coordinatedto.DTO, []unitdto.DTO, error) {
	gameRoom, err := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return nil, nil, err
	}

	coordinates, err := coordinatedto.FromDTOList(coordinateDTOs)
	if err != nil {
		return nil, nil, err
	}

	area, err := areadto.FromDTO(areaDTO)
	if err != nil {
		return nil, nil, err
	}

	coordinatesInArea := area.FilterCoordinates(coordinates)
	units, err := gameRoom.GetUnitsWithCoordinates(coordinatesInArea)
	if err != nil {
		return nil, nil, err
	}

	return coordinatedto.ToDTOList(coordinatesInArea), unitdto.ToDTOList(units), nil
}

func (grs *serviceImplement) ReviveUnits(gameId uuid.UUID, coordinateDTOs []coordinatedto.DTO) error {
	if grs.gameUnitsRevivedEvent == nil {
		return ErrEventNotFound
	}

	coordinates, err := coordinatedto.FromDTOList(coordinateDTOs)
	if err != nil {
		return err
	}

	revivedAt, err := grs.gameRoomDomainService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return err
	}

	grs.gameUnitsRevivedEvent.Publish(gameId, coordinateDTOs, revivedAt)

	return nil
}
