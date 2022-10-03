package memoryrepository

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/game/repository"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound       = errors.New("the game room with the id was not found")
	ErrGameRoomLockerNotFound = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists    = errors.New("the player with the given id alredy exists in the game room")
)

type gameRoomRepositoryMemory struct {
	records       map[uuid.UUID]*aggregate.GameRoom
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var gameRoomRepositoryMemoryInstance *gameRoomRepositoryMemory

func GetGameRoomRepositoryMemory() repository.GameRoomRepository {
	if gameRoomRepositoryMemoryInstance == nil {
		gameRoomRepositoryMemoryInstance = &gameRoomRepositoryMemory{
			records:       make(map[uuid.UUID]*aggregate.GameRoom),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return gameRoomRepositoryMemoryInstance
	} else {
		return gameRoomRepositoryMemoryInstance
	}
}

func (m *gameRoomRepositoryMemory) Get(id uuid.UUID) (aggregate.GameRoom, error) {
	record, exists := m.records[id]
	if !exists {
		return aggregate.GameRoom{}, ErrGameRoomNotFound
	}

	return *record, nil
}

func (m *gameRoomRepositoryMemory) Update(id uuid.UUID, gameRoom aggregate.GameRoom) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameRoomNotFound
	}

	m.records[id] = &gameRoom

	return nil
}

func (m *gameRoomRepositoryMemory) GetAll() []aggregate.GameRoom {
	gameRooms := make([]aggregate.GameRoom, 0)
	for _, record := range m.records {
		gameRooms = append(gameRooms, *record)
	}
	return gameRooms
}

func (m *gameRoomRepositoryMemory) Add(gameRoom aggregate.GameRoom) error {
	m.records[gameRoom.GetGameId()] = &gameRoom
	m.recordLockers[gameRoom.GetGameId()] = &sync.RWMutex{}

	return nil
}

func (m *gameRoomRepositoryMemory) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameRoomLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *gameRoomRepositoryMemory) LockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameRoomLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
