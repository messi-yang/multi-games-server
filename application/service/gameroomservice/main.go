package gameroomservice

import (
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitmapdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/event/gamecomputedevent"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type GameRoomService interface {
	CreateRoom(width int, height int) (gameId uuid.UUID, err error)
	GetUnitMapWithArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitmapdto.UnitMapDTO, error)
	TcikAllUnitMaps()
}

type gameRoomServiceImplement struct {
	gameRoomDomainService gameroomservice.GameRoomService
	gameComputeEvent      gamecomputedevent.GameComputedEvent
	locker                sync.RWMutex
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

func (grs *gameRoomServiceImplement) GetUnitMapWithArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitmapdto.UnitMapDTO, error) {
	area, err := areadto.FromDTO(areaDTO)
	if err != nil {
		return unitmapdto.UnitMapDTO{}, err
	}
	unitMap, err := grs.gameRoomDomainService.GetUnitMapWithArea(gameId, area)
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
