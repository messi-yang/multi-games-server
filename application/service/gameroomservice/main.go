package gameroomservice

import (
	"errors"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitmapdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitmapupdatedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gameunitsrevivedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

var (
	ErrEventNotFound = errors.New("event was not found")
)

type GameRoomService interface {
	CreateRoom(width int, height int) (gameId uuid.UUID, err error)
	GetUnitMapByArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitMapDTO unitmapdto.UnitMapDTO, receivedAt time.Time, err error)
	TcikAllUnitMaps() error
	ReviveUnits(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) error
	GetUnitMapSize(gameId uuid.UUID) (mapsizedto.MapSizeDTO, error)
	GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, areaDTO areadto.AreaDTO) ([]coordinatedto.CoordinateDTO, []unitdto.UnitDTO, error)
}

type gameRoomServiceImplement struct {
	gameRoomDomainService   gameroomservice.GameRoomService
	gameUnitMapUpdatedEvent gameunitmapupdatedevent.GameUnitMapUpdatedEvent
	gameUnitsRevivedEvent   gameunitsrevivedevent.UnitsRevivedEvent
}

type Configuration struct {
	GameRoomRepository gameroomrepository.GameRoomRepository
	GameComputeEvent   gameunitmapupdatedevent.GameUnitMapUpdatedEvent
	UnitsRevivedEvent  gameunitsrevivedevent.UnitsRevivedEvent
}

func NewGameRoomService(config Configuration) GameRoomService {
	return &gameRoomServiceImplement{
		gameRoomDomainService:   gameroomservice.NewGameRoomService(config.GameRoomRepository),
		gameUnitMapUpdatedEvent: config.GameComputeEvent,
		gameUnitsRevivedEvent:   config.UnitsRevivedEvent,
	}
}

func (grs *gameRoomServiceImplement) CreateRoom(width int, height int) (gameId uuid.UUID, err error) {
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

func (grs *gameRoomServiceImplement) GetUnitMapByArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitmapdto.UnitMapDTO, time.Time, error) {
	area, err := areadto.FromDTO(areaDTO)
	if err != nil {
		return unitmapdto.UnitMapDTO{}, time.Time{}, err
	}
	unitMap, receivedAt, err := grs.gameRoomDomainService.GetUnitMapByArea(gameId, area)
	if err != nil {
		return unitmapdto.UnitMapDTO{}, time.Time{}, err
	}

	unitMapDTO := unitmapdto.ToDTO(unitMap)

	return unitMapDTO, receivedAt, nil
}

func (grs *gameRoomServiceImplement) TcikAllUnitMaps() error {
	if grs.gameUnitMapUpdatedEvent == nil {
		return ErrEventNotFound
	}

	gameRooms := grs.gameRoomDomainService.GetAllRooms()
	for _, gameRoom := range gameRooms {
		updatedAt, err := grs.gameRoomDomainService.TickUnitMap(gameRoom.GetGameId())
		if err != nil {
			continue
		}
		grs.gameUnitMapUpdatedEvent.Publish(gameRoom.GetGameId(), updatedAt)
	}

	return nil
}

func (grs *gameRoomServiceImplement) GetUnitMapSize(gameId uuid.UUID) (mapsizedto.MapSizeDTO, error) {
	gameRoom, err := grs.gameRoomDomainService.GetRoom(gameId)
	if err != nil {
		return mapsizedto.MapSizeDTO{}, err
	}
	return mapsizedto.ToDTO(gameRoom.GetUnitMapSize()), nil
}

func (grs *gameRoomServiceImplement) GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, areaDTO areadto.AreaDTO) ([]coordinatedto.CoordinateDTO, []unitdto.UnitDTO, error) {
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

func (grs *gameRoomServiceImplement) ReviveUnits(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) error {
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
