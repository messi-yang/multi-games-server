package gameroomservice

import (
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/coordinatedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/mapsizedto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitmapdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/coordinatesupdatedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameRoomService interface {
	CreateRoom(width int, height int) (gameId uuid.UUID, err error)
	GetUnitMapByArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitmapdto.UnitMapDTO, error)
	TcikAllUnitMaps()
	ReviveUnits(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO) error
	GetUnitMapSize(gameId uuid.UUID) (mapsizedto.MapSizeDTO, error)
	GetUnitsByCoordinatesInArea(gameId uuid.UUID, coordinateDTOs []coordinatedto.CoordinateDTO, areaDTO areadto.AreaDTO) ([]coordinatedto.CoordinateDTO, []unitdto.UnitDTO, error)
}

type gameRoomServiceImplement struct {
	gameRoomDomainService   gameroomservice.GameRoomService
	gameComputeEvent        gamecomputedevent.GameComputedEvent
	coordinatesUpdatedEvent coordinatesupdatedevent.CoordinatesUpdatedEvent
	locker                  sync.RWMutex
}

func NewGameRoomService(gameRoomRepository gameroomrepository.GameRoomRepository) GameRoomService {
	return &gameRoomServiceImplement{
		gameRoomDomainService: gameroomservice.NewGameRoomService(gameRoomRepository),
		locker:                sync.RWMutex{},
	}
}

func NewGameRoomServiceWithGameComputedEvent(gameRoomRepository gameroomrepository.GameRoomRepository, gameComputeEvent gamecomputedevent.GameComputedEvent) GameRoomService {
	return &gameRoomServiceImplement{
		gameRoomDomainService: gameroomservice.NewGameRoomService(gameRoomRepository),
		gameComputeEvent:      gameComputeEvent,
		locker:                sync.RWMutex{},
	}
}

func NewGameRoomServiceWithCoordinatesUpdatedEvent(gameRoomRepository gameroomrepository.GameRoomRepository, coordinatesUpdatedEvent coordinatesupdatedevent.CoordinatesUpdatedEvent) GameRoomService {
	return &gameRoomServiceImplement{
		gameRoomDomainService:   gameroomservice.NewGameRoomService(gameRoomRepository),
		coordinatesUpdatedEvent: coordinatesUpdatedEvent,
		locker:                  sync.RWMutex{},
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

func (grs *gameRoomServiceImplement) GetUnitMapByArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitmapdto.UnitMapDTO, error) {
	area, err := areadto.FromDTO(areaDTO)
	if err != nil {
		return unitmapdto.UnitMapDTO{}, err
	}
	unitMap, err := grs.gameRoomDomainService.GetUnitMapByArea(gameId, area)
	if err != nil {
		return unitmapdto.UnitMapDTO{}, err
	}

	unitMapDTO := unitmapdto.ToDTO(unitMap)

	return unitMapDTO, nil
}

func (grs *gameRoomServiceImplement) TcikAllUnitMaps() {
	gameRooms := grs.gameRoomDomainService.GetAllRooms()
	for _, gameRoom := range gameRooms {
		grs.gameRoomDomainService.TickUnitMap(gameRoom.GetGameId())
		grs.gameComputeEvent.Publish(gameRoom.GetGameId())
	}
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
	coordinates, err := coordinatedto.FromDTOList(coordinateDTOs)
	if err != nil {
		return err
	}

	err = grs.gameRoomDomainService.ReviveUnits(gameId, coordinates)
	if err != nil {
		return err
	}

	grs.coordinatesUpdatedEvent.Publish(gameId, coordinateDTOs)

	return nil
}
