package gameroommemory

import (
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository/gameroomrepository"
	"github.com/google/uuid"
)

type memory struct {
	records       map[uuid.UUID]*aggregate.GameRoom
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var memoryInstance *memory

func GetRepository() gameroomrepository.Repository {
	if memoryInstance == nil {
		memoryInstance = &memory{
			records:       make(map[uuid.UUID]*aggregate.GameRoom),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return memoryInstance
	} else {
		return memoryInstance
	}
}

func (m *memory) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	record, exists := m.records[id]
	if !exists {
		return aggregate.GameRoom{}, gameroomrepository.ErrGameRoomNotFound
	}

	return *record, nil
}

func (m *memory) Update(id uuid.UUID, gameRoom aggregate.GameRoom) error {
	_, exists := m.records[id]
	if !exists {
		return gameroomrepository.ErrGameRoomNotFound
	}

	m.records[id] = &gameRoom

	return nil
}

func (m *memory) GetAll() []aggregate.GameRoom {
	gameRooms := make([]aggregate.GameRoom, 0)
	for _, record := range m.records {
		gameRooms = append(gameRooms, *record)
	}
	return gameRooms
}

func (m *memory) Add(gameRoom aggregate.GameRoom) error {
	m.records[gameRoom.GetGameId()] = &gameRoom
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
