package memoryrepository

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/google/uuid"
)

var (
	ErrGameRoomNotFound       = errors.New("the game room with the id was not found")
	ErrGameRoomLockerNotFound = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists    = errors.New("the player with the given id alredy exists in the game room")
)

type gameRoomMemoryRepository struct {
	records       map[uuid.UUID]*game.GameRoom
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var gameRoomMemoryRepositoryInstance *gameRoomMemoryRepository

func NewGameRoomMemoryRepository() game.GameRoomRepository {
	if gameRoomMemoryRepositoryInstance == nil {
		gameRoomMemoryRepositoryInstance = &gameRoomMemoryRepository{
			records:       make(map[uuid.UUID]*game.GameRoom),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return gameRoomMemoryRepositoryInstance
	}
	return gameRoomMemoryRepositoryInstance
}

func (m *gameRoomMemoryRepository) Get(id uuid.UUID) (game.GameRoom, error) {
	record, exists := m.records[id]
	if !exists {
		return game.GameRoom{}, ErrGameRoomNotFound
	}

	return *record, nil
}

func (m *gameRoomMemoryRepository) Update(id uuid.UUID, gameRoom game.GameRoom) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameRoomNotFound
	}

	m.records[id] = &gameRoom

	return nil
}

func (m *gameRoomMemoryRepository) GetAll() []game.GameRoom {
	gameRooms := make([]game.GameRoom, 0)
	for _, record := range m.records {
		gameRooms = append(gameRooms, *record)
	}
	return gameRooms
}

func (m *gameRoomMemoryRepository) Add(gameRoom game.GameRoom) error {
	m.records[gameRoom.GetId()] = &gameRoom
	m.recordLockers[gameRoom.GetId()] = &sync.RWMutex{}

	return nil
}

func (m *gameRoomMemoryRepository) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameRoomLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *gameRoomMemoryRepository) LockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameRoomLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
