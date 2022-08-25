package gameroommemory

import (
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type gameRoomRecord struct {
	unitMap valueobject.UnitMap
}

type gameRoomMemory struct {
	gameRoomRecords map[uuid.UUID]gameRoomRecord
	locker          sync.RWMutex
}

var gameRoomMemoryInstance *gameRoomMemory

func GetGameRoomMemory() gameroomrepository.GameRoomRepository {
	if gameRoomMemoryInstance == nil {
		gameRoomMemoryInstance = &gameRoomMemory{
			gameRoomRecords: make(map[uuid.UUID]gameRoomRecord),
			locker:          sync.RWMutex{},
		}
		return gameRoomMemoryInstance
	} else {
		return gameRoomMemoryInstance
	}
}

func (gmi *gameRoomMemory) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	gmi.locker.RLock()
	defer gmi.locker.RUnlock()

	gameRoomRecord, exists := gmi.gameRoomRecords[id]
	if !exists {
		return aggregate.GameRoom{}, gameroomrepository.ErrGameRoomNotFound
	}

	game := entity.NewGameFromExistingEntity(id, gameRoomRecord.unitMap)
	gameRoom := aggregate.NewGameRoom(game)
	return gameRoom, nil
}

func (gmi *gameRoomMemory) GetAll() []aggregate.GameRoom {
	gmi.locker.RLock()
	defer gmi.locker.RUnlock()

	gameRoom := make([]aggregate.GameRoom, 0)
	for gameId, gameRoomRecord := range gmi.gameRoomRecords {
		game := entity.NewGameFromExistingEntity(gameId, gameRoomRecord.unitMap)
		newGameRoom := aggregate.NewGameRoom(game)
		gameRoom = append(gameRoom, newGameRoom)
	}
	return gameRoom
}

func (gmi *gameRoomMemory) UpdateUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate, units []valueobject.Unit) error {
	gmi.locker.Lock()
	defer gmi.locker.Unlock()

	gameRoomRecord, exists := gmi.gameRoomRecords[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	for coordIdx, coord := range coordinates {
		gameRoomRecord.unitMap.SetUnit(coord, units[coordIdx])
	}

	return nil
}

func (gmi *gameRoomMemory) UpdateUnitMap(gameId uuid.UUID, unitMap valueobject.UnitMap) error {
	gmi.locker.Lock()
	defer gmi.locker.Unlock()

	gameRoomRecord, exists := gmi.gameRoomRecords[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	gameRoomRecord.unitMap = unitMap

	return nil
}

func (gmi *gameRoomMemory) Add(gameRoom aggregate.GameRoom) error {
	gameUnitMap := gameRoom.GetUnitMap()
	gmi.gameRoomRecords[gameRoom.GetGameId()] = gameRoomRecord{
		unitMap: gameUnitMap,
	}

	return nil
}
