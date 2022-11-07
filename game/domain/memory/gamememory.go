package memory

import (
	"errors"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/aggregate"
	"github.com/dum-dum-genius/game-of-liberty-computer/game/domain/repository"
	"github.com/google/uuid"
)

var (
	ErrGameNotFound        = errors.New("the game room with the id was not found")
	ErrGameLockerNotFound  = errors.New("the game room locker for with the id was not found")
	ErrPlayerAlreadyExists = errors.New("the player with the given id alredy exists in the game room")
)

type gameMemory struct {
	records       map[uuid.UUID]*aggregate.Game
	recordLockers map[uuid.UUID]*sync.RWMutex
}

var gameMemoryInstance *gameMemory

func NewGameMemory() repository.GameRepository {
	if gameMemoryInstance == nil {
		gameMemoryInstance = &gameMemory{
			records:       make(map[uuid.UUID]*aggregate.Game),
			recordLockers: make(map[uuid.UUID]*sync.RWMutex),
		}
		return gameMemoryInstance
	}
	return gameMemoryInstance
}

func (m *gameMemory) Get(id uuid.UUID) (aggregate.Game, error) {
	record, exists := m.records[id]
	if !exists {
		return aggregate.Game{}, ErrGameNotFound
	}

	return *record, nil
}

func (m *gameMemory) Update(id uuid.UUID, game aggregate.Game) error {
	_, exists := m.records[id]
	if !exists {
		return ErrGameNotFound
	}

	m.records[id] = &game

	return nil
}

func (m *gameMemory) GetAll() []aggregate.Game {
	games := make([]aggregate.Game, 0)
	for _, record := range m.records {
		games = append(games, *record)
	}
	return games
}

func (m *gameMemory) Add(game aggregate.Game) error {
	m.records[game.GetId()] = &game
	m.recordLockers[game.GetId()] = &sync.RWMutex{}

	return nil
}

func (m *gameMemory) ReadLockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.RLock()
	return recordLocker.RUnlock, nil
}

func (m *gameMemory) LockAccess(gameId uuid.UUID) (func(), error) {
	recordLocker, exists := m.recordLockers[gameId]
	if !exists {
		return nil, ErrGameLockerNotFound
	}

	recordLocker.Lock()
	return recordLocker.Unlock, nil
}
