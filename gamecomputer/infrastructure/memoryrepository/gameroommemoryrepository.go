package memoryrepository

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/shared/domain/model/game"
	"github.com/google/uuid"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type gameMemoryRepository struct {
	records       map[uuid.UUID]*game.Game
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var gameMemoryRepositoryInstance *gameMemoryRepository

func NewGameMemoryRepository() game.GameRepository {
	if gameMemoryRepositoryInstance == nil {
		gameMemoryRepositoryInstance = &gameMemoryRepository{
			records:       make(map[uuid.UUID]*game.Game),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return gameMemoryRepositoryInstance
	}
	return gameMemoryRepositoryInstance
}

func (m *gameMemoryRepository) Get(id uuid.UUID) (game.Game, error) {
	record, exists := m.records[id]
	if !exists {
		return game.Game{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *gameMemoryRepository) Update(id uuid.UUID, game game.Game) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &game

	return nil
}

func (m *gameMemoryRepository) GetAll() []game.Game {
	games := make([]game.Game, 0)
	for _, record := range m.records {
		games = append(games, *record)
	}
	return games
}

func (m *gameMemoryRepository) Add(game game.Game) error {
	m.records[game.GetId()] = &game
	m.recordLockers[game.GetId()] = &sync.RWMutex{}

	return nil
}

func (m *gameMemoryRepository) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *gameMemoryRepository) LockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
