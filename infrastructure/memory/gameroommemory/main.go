package gameroommemory

import (
	"sync"
	"time"

	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/entity"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/repository/gameroomrepository"
	"github.com/dum-dum-genius/game-of-liberty-computer/domain/game/valueobject"
	"github.com/google/uuid"
)

type gameRoomRecord struct {
	unitMap  valueobject.UnitMap
	tickedAt time.Time
}

type gameRoomMemory struct {
	gameRoomRecords       map[uuid.UUID]*gameRoomRecord
	gameRoomRecordLockers map[uuid.UUID]*sync.RWMutex
}

var gameRoomMemoryInstance *gameRoomMemory

func GetGameRoomMemory() gameroomrepository.GameRoomRepository {
	if gameRoomMemoryInstance == nil {
		gameRoomMemoryInstance = &gameRoomMemory{
			gameRoomRecords:       make(map[uuid.UUID]*gameRoomRecord),
			gameRoomRecordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return gameRoomMemoryInstance
	} else {
		return gameRoomMemoryInstance
	}
}

func (gmi *gameRoomMemory) Get(id uuid.UUID) (aggregate.GameRoom, time.Time, error) {
	gameRoomRecord, exists := gmi.gameRoomRecords[id]
	if !exists {
		return aggregate.GameRoom{}, time.Time{}, gameroomrepository.ErrGameRoomNotFound
	}

	game := entity.NewGameFromExistingEntity(id, gameRoomRecord.unitMap)
	gameRoom := aggregate.NewGameRoomWithLastTickedAt(game, gameRoomRecord.tickedAt)
	return gameRoom, time.Now(), nil
}

func (gmi *gameRoomMemory) GetAll() []aggregate.GameRoom {
	gameRoom := make([]aggregate.GameRoom, 0)
	for gameId, gameRoomRecord := range gmi.gameRoomRecords {
		game := entity.NewGameFromExistingEntity(gameId, gameRoomRecord.unitMap)
		newGameRoom := aggregate.NewGameRoomWithLastTickedAt(game, gameRoomRecord.tickedAt)
		gameRoom = append(gameRoom, newGameRoom)
	}
	return gameRoom
}

func (gmi *gameRoomMemory) GetLastTickedAt(id uuid.UUID) (time.Time, error) {
	gameRoomRecord, exists := gmi.gameRoomRecords[id]
	if !exists {
		return time.Time{}, gameroomrepository.ErrGameRoomNotFound
	}

	return gameRoomRecord.tickedAt, nil
}

func (gmi *gameRoomMemory) UpdateUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate, units []valueobject.Unit) error {
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
	gameRoomRecord, exists := gmi.gameRoomRecords[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	gameRoomRecord.unitMap = unitMap

	return nil
}

func (gmi *gameRoomMemory) UpdateTickedAt(id uuid.UUID, tickedAt time.Time) error {
	gameRoomRecord, exists := gmi.gameRoomRecords[id]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	gameRoomRecord.tickedAt = tickedAt

	return nil
}

func (gmi *gameRoomMemory) Add(gameRoom aggregate.GameRoom) error {
	gameUnitMap := gameRoom.GetUnitMap()
	gmi.gameRoomRecords[gameRoom.GetGameId()] = &gameRoomRecord{
		unitMap: gameUnitMap,
	}
	gmi.gameRoomRecordLockers[gameRoom.GetGameId()] = &sync.RWMutex{}

	return nil
}

func (gmi *gameRoomMemory) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	gameRoomRecordLocker, exists := gmi.gameRoomRecordLockers[gameId]
	if !exists {
		return nil, gameroomrepository.ErrGameRoomLockerNotFound
	}

	gameRoomRecordLocker.RLock()
	return gameRoomRecordLocker.RUnlock, nil
}

func (gmi *gameRoomMemory) LockAccess(gameId uuid.UUID) (func(), error) {
	gameRoomRecordLocker, exists := gmi.gameRoomRecordLockers[gameId]
	if !exists {
		return nil, gameroomrepository.ErrGameRoomLockerNotFound
	}

	gameRoomRecordLocker.Lock()
	return gameRoomRecordLocker.Unlock, nil
}
