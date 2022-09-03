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

type record struct {
	unitMap  valueobject.UnitMap
	tickedAt time.Time
}

type memory struct {
	records       map[uuid.UUID]*record
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var memoryInstance *memory

func GetRepository() gameroomrepository.Repository {
	if memoryInstance == nil {
		memoryInstance = &memory{
			records:       make(map[uuid.UUID]*record),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return memoryInstance
	} else {
		return memoryInstance
	}
}

func (m *memory) Get(id uuid.UUID) (aggregate.GameRoom, time.Time, error) {
	record, exists := m.records[id]
	if !exists {
		return aggregate.GameRoom{}, time.Time{}, gameroomrepository.ErrGameRoomNotFound
	}

	game := entity.NewGameFromExistingEntity(id, record.unitMap)
	gameRoom := aggregate.NewGameRoomWithLastTickedAt(game, record.tickedAt)
	return gameRoom, time.Now(), nil
}

func (m *memory) GetAll() []aggregate.GameRoom {
	gameRoom := make([]aggregate.GameRoom, 0)
	for gameId, record := range m.records {
		game := entity.NewGameFromExistingEntity(gameId, record.unitMap)
		newGameRoom := aggregate.NewGameRoomWithLastTickedAt(game, record.tickedAt)
		gameRoom = append(gameRoom, newGameRoom)
	}
	return gameRoom
}

func (m *memory) GetLastTickedAt(id uuid.UUID) (time.Time, error) {
	record, exists := m.records[id]
	if !exists {
		return time.Time{}, gameroomrepository.ErrGameRoomNotFound
	}

	return record.tickedAt, nil
}

func (m *memory) UpdateUnits(gameId uuid.UUID, coordinates []valueobject.Coordinate, units []valueobject.Unit) error {
	record, exists := m.records[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	for coordIdx, coord := range coordinates {
		record.unitMap.SetUnit(coord, units[coordIdx])
	}

	return nil
}

func (m *memory) UpdateUnitMap(gameId uuid.UUID, unitMap valueobject.UnitMap) error {
	record, exists := m.records[gameId]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	record.unitMap = unitMap

	return nil
}

func (m *memory) UpdateLastTickedAt(id uuid.UUID, tickedAt time.Time) error {
	record, exists := m.records[id]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	record.tickedAt = tickedAt

	return nil
}

func (m *memory) Add(gameRoom aggregate.GameRoom) error {
	gameUnitMap := gameRoom.GetUnitMap()
	m.records[gameRoom.GetGameId()] = &record{
		unitMap: gameUnitMap,
	}
	m.recordLockers[gameRoom.GetGameId()] = &sync.RWMutex{}

	return nil
}

func (m *memory) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, gameroomrepository.ErrGameRoomLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *memory) LockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, gameroomrepository.ErrGameRoomLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
