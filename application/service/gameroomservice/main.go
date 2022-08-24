package gameroomservice

import (
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/areadto"
	"github.com/dum-dum-genius/game-of-liberty-computer/application/dto/unitmapdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/service/gameroomservice"
	"github.com/google/uuid"
)

type GameRoomService interface {
	GetUnitMapWithArea(gameId uuid.UUID, areaDTO areadto.AreaDTO) (unitmapdto.UnitMapDTO, error)
}

type gameRoomServiceImplement struct {
	gameRoomDomainService gameroomservice.GameRoomService
	locker                sync.RWMutex
}

var gameRoomService GameRoomService = nil

func NewGameRoomService(gameRoomRepository gameroomrepository.GameRoomRepository) GameRoomService {
	if gameRoomService == nil {
		gameRoomService = &gameRoomServiceImplement{
			gameRoomDomainService: gameroomservice.NewGameRoomService(gameRoomRepository),
			locker:                sync.RWMutex{},
		}
	}
	return gameRoomService
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
